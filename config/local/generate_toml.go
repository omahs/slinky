//go:build ignore
// +build ignore

package main

import (
	"flag"
	"fmt"
	"github.com/skip-mev/slinky/providers/apis/uniswap"
	"os"
	"time"

	"github.com/BurntSushi/toml"

	"github.com/skip-mev/slinky/oracle/config"
	slinkytypes "github.com/skip-mev/slinky/pkg/types"
)

var (
	// oracleCfgPath is the path to write the oracle config file to. By default, this
	// will write the oracle config file to the local directory.
	oracleCfgPath = flag.String("oracle-config-path", "oracle.toml", "path to write the oracle config file to")

	// LocalConfig defines a readable config for local development. Any changes to this
	// file should be reflected in oracle.toml. To update the oracle.toml file, run
	// `make update-local-config`. This will update any changes to the oracle.toml file
	// as they are made to this file.
	LocalConfig = config.OracleConfig{
		// -----------------------------------------------------------	//
		// -----------------Aggregate Market Config-------------------	//
		// -----------------------------------------------------------	//
		Market: config.AggregateMarketConfig{
			Feeds:           Feeds,
			AggregatedFeeds: AggregatedFeeds,
		},
		Production: true,
		// -----------------------------------------------------------	//
		// ----------------------Metrics Config-----------------------	//
		// -----------------------------------------------------------	//
		Metrics: config.MetricsConfig{
			Enabled:                 true,
			PrometheusServerAddress: "0.0.0.0:8002",
		},
		UpdateInterval: 1500 * time.Millisecond,
		Providers: []config.ProviderConfig{
			// -----------------------------------------------------------	//
			// ---------------------Start API Providers--------------------	//
			// -----------------------------------------------------------	//
			//
			// NOTE: Some of the provider's are only capable of fetching data for a subset of
			// all currency pairs. Before adding a new market to the oracle, ensure that
			// the provider supports fetching data for the currency pair.
			// -----------------------------------------------------------	//
			// ---------------------Start WebSocket Providers--------------	//
			// -----------------------------------------------------------	//
			//
			// NOTE: Some of the provider's are only capable of fetching data for a subset of
			// all currency pairs. Before adding a new market to the oracle, ensure that
			// the provider supports fetching data for the currency pair.
			{
				Name:   uniswap.Name,
				API:    uniswap.DefaultAPIConfig,
				Market: uniswap.DefaultMarketConfig,
			},
		},
	}

	// Feeds is a map of all of the price feeds that the oracle will fetch prices for.
	Feeds = map[string]config.FeedConfig{
		"WETH/USDC": {
			CurrencyPair: slinkytypes.NewCurrencyPair("WETH", "USDC"),
		},
	}

	// AggregatedFeeds is a map of all of the conversion markets that will be used to convert
	// all of the price feeds into a common set of currency pairs.
	AggregatedFeeds = map[string]config.AggregateFeedConfig{
		"WETH/USDC": {
			CurrencyPair: slinkytypes.NewCurrencyPair("WETH", "USDC"),
			Conversions: []config.Conversions{
				{
					{
						CurrencyPair: slinkytypes.NewCurrencyPair("WETH", "USDC"),
						Invert:       false,
					},
				},
			},
		},
	}
)

// main executes a simple script that encodes the local config file to the local
// directory.
func main() {
	flag.Parse()

	// Open the local config file. This will overwrite any changes made to the
	// local config file.
	f, err := os.Create(*oracleCfgPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error creating local config file: %v\n", err)
	}
	defer f.Close()

	if err := LocalConfig.ValidateBasic(); err != nil {
		fmt.Fprintf(os.Stderr, "error validating local config: %v\n", err)
		return
	}

	// Encode the local config file.
	encoder := toml.NewEncoder(f)
	if err := encoder.Encode(LocalConfig); err != nil {
		panic(err)
	}
}
