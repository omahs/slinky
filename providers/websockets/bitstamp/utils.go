package bitstamp

import (
	"time"

	"github.com/skip-mev/slinky/oracle/config"
	"github.com/skip-mev/slinky/oracle/types"
)

const (
	// Name is the name of the bitstamp provider.
	Name = "bitstamp_ws"

	Type = types.ConfigType

	// WSS is the bitstamp websocket address.
	WSS = "wss://ws.bitstamp.net"

	// DefaultPingInterval is the default ping interval for the bitstamp websocket.
	DefaultPingInterval = 10 * time.Second
)

var (
	// DefaultWebSocketConfig returns the default websocket config for bitstamp.
	DefaultWebSocketConfig = config.WebSocketConfig{
		Enabled:                       true,
		Name:                          Name,
		MaxBufferSize:                 config.DefaultMaxBufferSize,
		ReconnectionTimeout:           config.DefaultReconnectionTimeout,
		WSS:                           WSS,
		ReadBufferSize:                config.DefaultReadBufferSize,
		WriteBufferSize:               config.DefaultWriteBufferSize,
		HandshakeTimeout:              config.DefaultHandshakeTimeout,
		EnableCompression:             config.DefaultEnableCompression,
		WriteTimeout:                  config.DefaultWriteTimeout,
		ReadTimeout:                   config.DefaultReadTimeout * 5,
		PingInterval:                  DefaultPingInterval,
		MaxReadErrorCount:             config.DefaultMaxReadErrorCount,
		MaxSubscriptionsPerConnection: config.DefaultMaxSubscriptionsPerConnection,
	}

	DefaultProviderConfig = config.ProviderConfig{
		Name:      Name,
		WebSocket: DefaultWebSocketConfig,
		Type:      Type,
	}
)
