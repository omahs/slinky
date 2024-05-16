package cryptodotcom

import (
	"github.com/skip-mev/slinky/oracle/config"
	"github.com/skip-mev/slinky/oracle/types"
)

const (
	// URL is the URL used to connect to the Crypto.com websocket API. This can be found here
	// https://exchange-docs.crypto.com/exchange/v1/rest-ws/index.html?javascript#websocket-root-endpoints
	// Note that Crypto.com offers a sandbox and production environment.

	// Name is the name of the Crypto.com provider.
	Name = "crypto_dot_com_ws"

	Type = types.ConfigType

	// URL_PROD is the URL used to connect to the Crypto.com production websocket API.
	URL_PROD = "wss://stream.crypto.com/exchange/v1/market"

	// URL_SANDBOX is the URL used to connect to the Crypto.com sandbox websocket API. This will
	// return static prices.
	URL_SANDBOX = "wss://uat-stream.3ona.co/exchange/v1/market"
)

var (
	// DefaultWebSocketConfig is the default configuration for the Crypto.com Websocket.
	DefaultWebSocketConfig = config.WebSocketConfig{
		Name:                          Name,
		Enabled:                       true,
		MaxBufferSize:                 config.DefaultMaxBufferSize,
		ReconnectionTimeout:           config.DefaultReconnectionTimeout,
		WSS:                           URL_PROD,
		ReadBufferSize:                config.DefaultReadBufferSize,
		WriteBufferSize:               config.DefaultWriteBufferSize,
		HandshakeTimeout:              config.DefaultHandshakeTimeout,
		EnableCompression:             config.DefaultEnableCompression,
		ReadTimeout:                   config.DefaultReadTimeout,
		WriteTimeout:                  config.DefaultWriteTimeout,
		PingInterval:                  config.DefaultPingInterval,
		MaxReadErrorCount:             config.DefaultMaxReadErrorCount,
		MaxSubscriptionsPerConnection: config.DefaultMaxSubscriptionsPerConnection,
	}

	DefaultProviderConfig = config.ProviderConfig{
		Name:      Name,
		WebSocket: DefaultWebSocketConfig,
		Type:      Type,
	}
)
