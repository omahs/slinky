package bitstamp

import (
	"time"

	"github.com/skip-mev/slinky/oracle/config"
	oracletypes "github.com/skip-mev/slinky/x/oracle/types"
)

const (
	// Name is the name of the bitstamp provider.
	Name = "bitstamp"

	// WSS is the bitstamp websocket address.
	WSS = "wss://ws.bitstamp.net"
)

var (
	// DefaultWebSocketConfig returns the default websocket config for bitstamp.
	DefaultWebSocketConfig = config.WebSocketConfig{
		Enabled:             true,
		Name:                Name,
		MaxBufferSize:       config.DefaultMaxBufferSize,
		ReconnectionTimeout: config.DefaultReconnectionTimeout,
		WSS:                 WSS,
		ReadBufferSize:      config.DefaultReadBufferSize,
		WriteBufferSize:     config.DefaultWriteBufferSize,
		HandshakeTimeout:    config.DefaultHandshakeTimeout,
		EnableCompression:   config.DefaultEnableCompression,
		WriteTimeout:        config.DefaultWriteTimeout,
		ReadTimeout:         config.DefaultReadTimeout,
		PingInterval:        10 * time.Second,
		MaxReadErrorCount:   config.DefaultMaxReadErrorCount,
	}

	// DefaultMarketConfig returns the default market config for bitstamp.
	DefaultMarketConfig = config.MarketConfig{
		Name: Name,
		CurrencyPairToMarketConfigs: map[string]config.CurrencyPairMarketConfig{
			"BITCOIN/USD/8": {
				Ticker:       "btcusd",
				CurrencyPair: oracletypes.NewCurrencyPair("BITCOIN", "USD", oracletypes.DefaultDecimals),
			},
			"ETHEREUM/USD/8": {
				Ticker:       "ethusd",
				CurrencyPair: oracletypes.NewCurrencyPair("ETHEREUM", "USD", oracletypes.DefaultDecimals),
			},
			"SOLANA/USD/8": {
				Ticker:       "solusd",
				CurrencyPair: oracletypes.NewCurrencyPair("SOLANA", "USD", oracletypes.DefaultDecimals),
			},
			"AVAX/USD/8": {
				Ticker:       "avaxusd",
				CurrencyPair: oracletypes.NewCurrencyPair("AVAX", "USD", oracletypes.DefaultDecimals),
			},
			"DYDX/USD/8": {
				Ticker:       "dydxusd",
				CurrencyPair: oracletypes.NewCurrencyPair("DYDX", "USD", oracletypes.DefaultDecimals),
			},
			"ETHEREUM/BITCOIN/8": {
				Ticker:       "ethbtc",
				CurrencyPair: oracletypes.NewCurrencyPair("ETHEREUM", "BITCOIN", oracletypes.DefaultDecimals),
			},
		},
	}
)
