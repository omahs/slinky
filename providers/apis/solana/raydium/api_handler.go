package raydium

import (
	"time"
	"encoding/json"
	"fmt"
	"math/big"
	"net/http"
	"github.com/gagliardetto/solana-go/programs/token"
	"github.com/skip-mev/slinky/pkg/math"

	bin "github.com/gagliardetto/binary"
	"github.com/gagliardetto/solana-go/rpc"
	"github.com/skip-mev/slinky/oracle/config"
	"github.com/skip-mev/slinky/oracle/types"
	oracletypes "github.com/skip-mev/slinky/oracle/types"
	slinkyjson "github.com/skip-mev/slinky/pkg/json"
	"github.com/skip-mev/slinky/providers/apis/solana/types/raydium"
	providertypes "github.com/skip-mev/slinky/providers/types"
	mmtypes "github.com/skip-mev/slinky/x/marketmap/types"
)

var _ oracletypes.PriceAPIDataHandler = (*APIHandler)(nil)

// APIHandler implements the APIDataHandlerWithBody interface for Raydium.
type APIHandler struct {
	// market is the config for the Raydium API.
	market types.ProviderMarketMap

	// api is the config for the Raydium API.
	api config.APIConfig

	// associatedAccountsPerTicker is the set of solana addresses that should be queried
	// for a given pool's account
	associatedAccountsPerTicker map[string]PoolAssociatedAccounts // ticker -> associated accounts
}

// NewAPIHandler returns a new Raydium APIDataHandlerWithBody. This method requires
// that the given market + config are valid, otherwise a nil implemnentation + error
// is returned.
func NewAPIHandler(
	market types.ProviderMarketMap,
	api config.APIConfig,
) (oracletypes.PriceAPIDataHandler, error) {
	if err := market.ValidateBasic(); err != nil {
		return nil, err
	}

	if market.Name != Name {
		return nil, fmt.Errorf("expected market config name raydium, got %s", market.Name)
	}

	if api.Name != Name {
		return nil, fmt.Errorf("expected api config name raydium, got %s", api.Name)
	}

	if !api.Enabled {
		return nil, fmt.Errorf("api config for raydium is not enabled")
	}

	if err := api.ValidateBasic(); err != nil {
		return nil, err
	}

	// generate the associated accounts per ticker
	associatedAccountsPerTicker := make(map[string]PoolAssociatedAccounts, len(market.OffChainMap))
	for _, ticker := range market.OffChainMap {
		tickerName := ticker.CurrencyPair.String()

		raydiumMetadata, err := unmarshalMetadataJSON(ticker.Metadata_JSON)
		if err != nil {
			return nil, fmt.Errorf("failed to unmarshal metadata for %s: %w", tickerName, err)
		}

		associatedAccounts, err := derivePoolAssociatedAccounts(raydiumDexProgramAccount, raydiumMetadata)
		if err != nil {
			return nil, fmt.Errorf("failed to derive associated accounts for %s: %w", tickerName, err)
		}

		associatedAccountsPerTicker[tickerName] = associatedAccounts
	}

	return &APIHandler{
		market: market,
		api:    api,
		associatedAccountsPerTicker: associatedAccountsPerTicker,
	}, nil
}

// CreateURL returns the URL to be queried for the Raydium API. Notice, the URL returned
// will be static, and should point to the JSON-RPC endpoint of a solana RPC-node.
func (h *APIHandler) CreateURL([]mmtypes.Ticker) (string, error) {
	return h.api.URL, nil
}

// CreateBody returns the body to be sent in the POST request to the Raydium API. The body
// returned will be designated for the solana getMultipleAccounts JSON-RPC method.
// 
// The body returned will be a standard getMultipleAccounts JSON-RPC request, with a body defined as follows:
//  - per ticker
//    - PoolAssociatedInfo account
//    - PoolQuoteTokenVault account
//    - PoolBaseTokenVault account
func (h *APIHandler) CreateBody(tickers []mmtypes.Ticker) ([]byte, error) {
	params := make([]interface{}, 3 * len(tickers) + 1)
	
	// populate the request params for each ticker
	for i, ticker := range tickers {
		tickerName := ticker.CurrencyPair.String()
		associatedAccounts, ok := h.associatedAccountsPerTicker[tickerName]
		if !ok {
			return nil, fmt.Errorf("missing associated accounts for ticker %s", tickerName)
		}
		
		// populate the associated accounts for the pool
		params[3 * i] = associatedAccounts.PoolAccountInfo
		params[3 * i + 1] = associatedAccounts.QuoteTokenPoolVaultAccount
		params[3 * i + 2] = associatedAccounts.BaseTokenPoolVaultAccount
	}
	// the first entry will be parameters for the JSON-RPC server
	params = append(params, map[string]interface{}{
		"commitment": rpc.CommitmentFinalized,
	})
	
	// create the JSON-RPC request
	req := slinkyjson.DefaultRPCRequest("getMultipleAccounts", params)

	return json.Marshal(req)
}

// ParseResponse parses the JSON response from the solana JSON-RPC request made in CreateBody.
// The response is expected to be a list of AccountInfo objects, where each object corresponds
// to the account information for a given pool.
func (h *APIHandler) ParseResponse(
	tickers []mmtypes.Ticker,
	resp *http.Response,
) providertypes.GetResponse[mmtypes.Ticker, *big.Int] {
	// unmarshal the HTTP response body
	defer resp.Body.Close()

	var rpcResp slinkyjson.RPCResponse
	if err := json.NewDecoder(resp.Body).Decode(&rpcResp); err != nil {
		return types.NewPriceResponseWithErr(
			tickers,
			providertypes.NewErrorWithCode(
				err,
				providertypes.ErrorInvalidResponse,
			),
		)
	}

	// check for errors in the RPC response
	if rpcResp.Error != nil {
		return types.NewPriceResponseWithErr(
			tickers,
			providertypes.NewErrorWithCode(
				rpcResp.Error,
				providertypes.ErrorInvalidResponse,
			),
		)
	}

	// get multiple accounts response
	var getMultipleAccountsResp rpc.GetMultipleAccountsResult
	if err := json.Unmarshal(rpcResp.Result, &getMultipleAccountsResp); err != nil {
		return types.NewPriceResponseWithErr(
			tickers,
			providertypes.NewErrorWithCode(
				err,
				providertypes.ErrorInvalidResponse,
			),
		)
	}

	var (
		resolved = make(types.ResolvedPrices)
		unresolved = make(types.UnResolvedPrices)
	)

	if len(getMultipleAccountsResp.Value) != 3 * len(tickers) + 1 {
		return types.NewPriceResponseWithErr(
			tickers,
			providertypes.NewErrorWithCode(
				fmt.Errorf("pool in response is missing"),
				providertypes.ErrorInvalidResponse,
			),
		)
	}

	// iterate over all accounts and attempt to parse into PoolAssociatedAccounts
	poolDataPerTicker := make(map[string]PoolAccountDataPerTicker)
	for i, account := range getMultipleAccountsResp.Value {
		switch i % 3 {
			// unmarshal as an AmmInfo
			case 0:
				var ammInfo raydium.AmmInfo
				if err := bin.NewBinDecoder(account.Data.GetBinary()).Decode(&ammInfo); err != nil {
					continue
				}

				// get the ticker at this index
				ticker := tickers[i / 3]

				_, ok := poolDataPerTicker[ticker.CurrencyPair.String()]
				if !ok {
					// initialize
					poolDataPerTicker[ticker.CurrencyPair.String()] = PoolAccountDataPerTicker{
						AmmInfo: &ammInfo,
					}
				} else {
					unresolved[ticker] = providertypes.UnresolvedResult{
						providertypes.NewErrorWithCode(
							fmt.Errorf("repeated associated accounts for ticker %s", ticker.CurrencyPair.String()),
							providertypes.ErrorInvalidResponse,
						),
					}
					continue
				}
			// unmarshal as the quote-token vault
			case 1:
				var tokenVault token.Account
				if err := bin.NewBinDecoder(account.Data.GetBinary()).Decode(&tokenVault); err != nil {
					continue
				}

				// get the ticker at this index
				ticker := tickers[i / 3]

				poolData, ok := poolDataPerTicker[ticker.CurrencyPair.String()]
				if !ok {
					unresolved[ticker] = providertypes.UnresolvedResult{
						providertypes.NewErrorWithCode(
							fmt.Errorf("missing associated accounts for ticker %s", ticker.CurrencyPair.String()),
							providertypes.ErrorInvalidResponse,
						),
					}
					continue
				}

				// if this ticker is already unresolved, skip
				if _, ok := unresolved[ticker]; ok {
					continue
				}

				poolData.QuoteTokenAccount = &tokenVault
				poolDataPerTicker[ticker.CurrencyPair.String()] = poolData
			// unmarshal as the base-token vault
			case 2:
				var tokenVault token.Account
				if err := bin.NewBinDecoder(account.Data.GetBinary()).Decode(&tokenVault); err != nil {
					continue
				}

				// get the ticker at this index
				ticker := tickers[i / 3]

				poolData, ok := poolDataPerTicker[ticker.CurrencyPair.String()]
				if !ok {
					unresolved[ticker] = providertypes.UnresolvedResult{
						providertypes.NewErrorWithCode(
							fmt.Errorf("missing associated accounts for ticker %s", ticker.CurrencyPair.String()),
							providertypes.ErrorInvalidResponse,
						),
					}
					continue
				}

				// if this ticker is already unresolved, skip
				if _, ok := unresolved[ticker]; ok {
					continue
				}

				poolData.BaseTokenAccount = &tokenVault
				poolDataPerTicker[ticker.CurrencyPair.String()] = poolData
		}
	}

	// iterate over all tickers and resolve the prices
	for _, ticker := range tickers {
		poolData, ok := poolDataPerTicker[ticker.CurrencyPair.String()]
		if !ok {
			unresolved[ticker] = providertypes.UnresolvedResult{
				providertypes.NewErrorWithCode(
					fmt.Errorf("missing associated accounts for ticker %s", ticker.CurrencyPair.String()),
					providertypes.ErrorInvalidResponse,
				),
			}
			continue
		}

		// if this ticker is already unresolved, skip
		if _, ok := unresolved[ticker]; ok {
			continue
		}

		// if any accounts are uninitialized, skip
		if poolData.AmmInfo == nil || poolData.BaseTokenAccount == nil || poolData.QuoteTokenAccount == nil {
			unresolved[ticker] = providertypes.UnresolvedResult{
				providertypes.NewErrorWithCode(
					fmt.Errorf("missing associated accounts for ticker %s", ticker.CurrencyPair.String()),
					providertypes.ErrorInvalidResponse,
				),
			}
			continue
		}

		ammInfo := *poolData.AmmInfo
		baseTokenAccount := *poolData.BaseTokenAccount
		quoteTokenAccount := *poolData.QuoteTokenAccount

		// determine exchange rate between base -> quote
		bigBaseTokens := new(big.Int).Mul(big.NewInt(int64(baseTokenAccount.Amount)), new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(18 - ammInfo.PcDecimals)), nil))
		bigQuoteTokens := new(big.Int).Mul(big.NewInt(int64(quoteTokenAccount.Amount)), new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(18 - ammInfo.CoinDecimals)), nil))

		unscaledPrice := new(big.Float).Quo(new(big.Float).SetInt(bigQuoteTokens), new(big.Float).SetInt(bigBaseTokens))

		// scale the price by decimals
		scaledPrice := math.BigFloatToBigInt(unscaledPrice, ticker.Decimals)
		resolved[ticker] = providertypes.NewResult(scaledPrice, time.Now())
	}

	return types.NewPriceResponse(resolved, unresolved)
}
