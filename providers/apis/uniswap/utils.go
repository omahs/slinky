package uniswap

import (
	"time"

	"github.com/skip-mev/slinky/oracle/config"
	slinkytypes "github.com/skip-mev/slinky/pkg/types"
)

// API does not require a subscription to use (i.e. No API key is required).

const (
	// Name is the name of the Binance provider.
	Name = "uniswap"
)

var (
	// DefaultAPIConfig is the default configuration for the Binance API.
	DefaultAPIConfig = config.APIConfig{
		Name:       Name,
		Atomic:     false,
		Enabled:    true,
		Timeout:    500 * time.Millisecond,
		Interval:   1 * time.Second,
		MaxQueries: 1,
		// TODO eric fixme
		URL: "https://eth.public-rpc.com/",
		// URL:        URLS,
		EVM: config.EVMAPIConfig{
			Enabled: true,
		},
	}

	// DefaultMarketConfig is the default US market configuration for Binance.
	DefaultMarketConfig = config.MarketConfig{
		Name: Name,
		CurrencyPairToMarketConfigs: map[string]config.CurrencyPairMarketConfig{
			"WETH/USDC": {
				Ticker:       "WETH/USDC",
				CurrencyPair: slinkytypes.NewCurrencyPair("WETH", "USDC"),
			},
		},
	}
)

type (
// Response is the expected response returned by the Binance API.
// The response is json formatted.
// Response format:
//
//	[
//  {
//    "symbol": "LTCBTC",
//    "price": "4.00000200"
//  },
//  {
//    "symbol": "ETHBTC",
//    "price": "0.07946600"
//  }
// ].
// Response []Data

// Data is the data returned by the Uniswap API.
// todo fixme based on the actual return values
/* Data struct {
   	Symbol string `json:"symbol"`
   	Price  string `json:"price"`
   }
*/
)
