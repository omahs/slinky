package uniswap

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	slinkytypes "github.com/skip-mev/slinky/pkg/types"
	"github.com/skip-mev/slinky/providers/apis/uniswap/factory"
	"github.com/skip-mev/slinky/providers/apis/uniswap/pool"
	handlers "github.com/skip-mev/slinky/providers/base/api/handlers/evm"
	providertypes "github.com/skip-mev/slinky/providers/types"
	"math/big"
	"time"
)

// UniswapRequestHandlerImpl is the default implementation of the RequestHandler interface.
type UniswapRequestHandlerImpl struct {
	// address is the contract's address on a chain corresponding to the provided RPC endpoint(s)
	address string
	// fee is the fee charged on the pool we want to fetch
	// Uniswap pools can be created with fees of 0.05%, 0.30%, or 1%
	// See https://docs.uniswap.org/concepts/protocol/fees
	// todo multiplex over all fee pools, for now we just hard code one
	fee *big.Int
	// url is an EVM JSON RPC endpoint
	url string
	// base and quote are hex addresses which correspond to the ERC20 token address for the asset
	// todo these should be stored in the mapping between canonical token pairs and provider-specific representations
	base, quote string
	// todo the amount you need to scale the returned price by is 10^base_decimals / 10^quote_decimals
	// ERC20 contracts each store their decimals in the chain--so we need to either get these from the chain or
	// just hard code them in the metadata field of the ticker config for each pair we're entering into the config
	scale int64
}

// NewUniswapRequestHandlerImpl creates a new RequestHandlerImpl. It manages making HTTP requests.
func NewUniswapRequestHandlerImpl(opts ...Option) (handlers.EVMRequestHandler[slinkytypes.CurrencyPair, *big.Int], error) {
	h := &UniswapRequestHandlerImpl{
		// todo I've hardcoded the fee -- different pools have different default fees
		// 3000 => 0.3%
		fee: big.NewInt(3000),
		// todo I've hardcoded the eth mainnet uniswap v3 contract address here
		address: "0x1F98431c8aD98523631AE4a59f267346ea31F984",
		// todo I've hardcoded the ethereum WETH ERC20 address here
		base: "0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2",
		// todo I've hardcoded the ethereum USDC ERC20 address here
		quote: "0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48",
		// todo I've hardcoded weth/usdc to scale by 1/10^12 (since WETH decimals is 18 and UDSC decimals is 6)
		scale: 12,
	}

	for _, opt := range opts {
		opt(h)
	}

	return h, nil
}

// todo need to refactor this to return some generic stuff instead of big.Int
func (r *UniswapRequestHandlerImpl) FetchPrices(ctx context.Context, url string, cps []slinkytypes.CurrencyPair) providertypes.GetResponse[slinkytypes.CurrencyPair, *big.Int] {
	if len(cps) != 1 {
		return providertypes.NewGetResponseWithErr[slinkytypes.CurrencyPair, *big.Int](
			cps,
			fmt.Errorf("expected 1 currency pair, got %d", len(cps)),
		)
	}
	cp := cps[0]
	client, err := ethclient.Dial(url)
	if err != nil {
		return providertypes.NewGetResponseWithErr[slinkytypes.CurrencyPair, *big.Int](cps, err)
	}
	uniFactory, err := factory.NewUniswap(common.HexToAddress(r.address), client)
	if err != nil {
		return providertypes.NewGetResponseWithErr[slinkytypes.CurrencyPair, *big.Int](cps, err)
	}

	// todo we might not need this bit--the contract addresses on mainnet are static so can be hardcoded in the config
	// also there seems to be a deterministic way to compute these https://github.com/daoleno/uniswapv3-sdk/blob/master/utils/compute_pool_address.go
	uniPoolAddr, err := uniFactory.GetPool(&bind.CallOpts{Context: ctx}, common.HexToAddress(r.base), common.HexToAddress(r.quote), r.fee)
	if err != nil {
		return providertypes.NewGetResponseWithErr[slinkytypes.CurrencyPair, *big.Int](cps, err)
	}
	uniPool, err := pool.NewUniswap(uniPoolAddr, client)
	if err != nil {
		return providertypes.NewGetResponseWithErr[slinkytypes.CurrencyPair, *big.Int](cps, err)
	}
	slotZero, err := uniPool.Slot0(&bind.CallOpts{Context: ctx})
	if err != nil {
		return providertypes.NewGetResponseWithErr[slinkytypes.CurrencyPair, *big.Int](cps, err)
	}
	// var sqrtPrice big.Int
	sqrtPrice, acc0 := slotZero.SqrtPriceX96.Float64()
	fmt.Printf("Lost %d digits of accuracy\n", acc0)
	bigSqrtPrice := big.NewFloat(sqrtPrice)
	var x96, acc1 = big.NewInt(0).Exp(big.NewInt(2), big.NewInt(96), nil).Float64()
	fmt.Printf("Lost %d digits of accuracy\n", acc1)
	var xScale, acc2 = big.NewInt(0).Exp(big.NewInt(10), big.NewInt(r.scale), nil).Float64()
	fmt.Printf("Lost %d digits of accuracy\n", acc2)
	bigX96 := big.NewFloat(x96)
	var priceFloat big.Float
	bigSqrtPrice.Quo(bigSqrtPrice, bigX96)
	priceFloat.Set(bigSqrtPrice.Mul(bigSqrtPrice, bigSqrtPrice))
	fmt.Printf("price float: %s\n", priceFloat.String())
	priceFloat.Quo(&priceFloat, big.NewFloat(xScale))
	priceFloat.Quo(big.NewFloat(1), &priceFloat)
	fmt.Printf("Actual price: %s\n", priceFloat.String())

	// todo this won't be right because it's unscaled--will lose precision and needs to be scaled by the number of decimals expected by slinky
	var price, _ = priceFloat.Int(nil)
	/*
		// scale the price based on the decimals of the two assets
		if r.scale > 0 {
			price.Mul(&price, big.NewInt(0).Exp(big.NewInt(10), big.NewInt(r.scale), nil))
		} else if r.scale < 0 {
			price.Mul(big.NewInt(0).Exp(big.NewInt(10), big.NewInt(r.scale), nil), &price)
		}
		fmt.Printf("Final price: %s\n", price.String())
		priceFloat, _ := price.Float64()
		fmt.Printf("Converts to: %f\n", 1/(priceFloat/100000000))
	*/
	// todo this price needs to be scaled by the decimals places of each of the input ERC20 contracts
	// todo and possible inverted based on whether the base/quote we desire matches with what is returned by uniswap
	// uniswap returns the price on the basis of sorted contract address
	return providertypes.NewGetResponse[slinkytypes.CurrencyPair, *big.Int](
		map[slinkytypes.CurrencyPair]providertypes.Result[*big.Int]{
			cp: providertypes.NewResult[*big.Int](price, time.Now()),
		},
		nil,
	)
}
