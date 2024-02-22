package oracle

import (
	"fmt"
	"github.com/skip-mev/slinky/oracle/config"
	slinkytypes "github.com/skip-mev/slinky/pkg/types"
	"github.com/skip-mev/slinky/providers/apis/uniswap"
	apihandlers "github.com/skip-mev/slinky/providers/base/api/handlers"
	handlers "github.com/skip-mev/slinky/providers/base/api/handlers/evm"
	"github.com/skip-mev/slinky/providers/base/api/metrics"
	"github.com/skip-mev/slinky/providers/types/factory"
	"go.uber.org/zap"
	"math/big"
)

// EVMAPIQueryHandlerFactory .
func EVMAPIQueryHandlerFactory() factory.APIQueryHandlerFactory[slinkytypes.CurrencyPair, *big.Int] {
	return func(logger *zap.Logger, cfg config.ProviderConfig, metrics metrics.APIMetrics) (apihandlers.APIQueryHandler[slinkytypes.CurrencyPair, *big.Int], error) {
		// Validate the provider config.
		err := cfg.ValidateBasic()
		if err != nil {
			return nil, err
		}

		var (
			requestHandler handlers.EVMRequestHandler[slinkytypes.CurrencyPair, *big.Int]
		)

		switch cfg.Name {
		case uniswap.Name:
			requestHandler, err = uniswap.NewUniswapRequestHandlerImpl()
			if err != nil {
				return nil, err
			}
		default:
			return nil, fmt.Errorf("unknown provider: %s", cfg.Name)
		}
		if err != nil {
			return nil, err
		}

		// Create the API query handler which encapsulates all the fetching and parsing logic.
		return handlers.NewEVMAPIQueryHandler[slinkytypes.CurrencyPair, *big.Int](
			logger,
			cfg.API,
			requestHandler,
			metrics,
		)
	}
}
