package handlers

import (
	"context"
	"fmt"
	"strings"
	"time"

	"golang.org/x/sync/errgroup"

	"go.uber.org/zap"

	"github.com/skip-mev/slinky/oracle/config"
	"github.com/skip-mev/slinky/providers/base/api/handlers"
	"github.com/skip-mev/slinky/providers/base/api/metrics"
	providertypes "github.com/skip-mev/slinky/providers/types"
)

// EVMAPIQueryHandlerImpl
type EVMAPIQueryHandlerImpl[K providertypes.ResponseKey, V providertypes.ResponseValue] struct {
	logger         *zap.Logger
	metrics        metrics.APIMetrics
	config         config.APIConfig
	requestHandler EVMRequestHandler[K, V]
}

// NewEVMAPIQueryHandler creates a new APIQueryHandler. It manages querying the data
// provider by using the APIDataHandler and RequestHandler.
func NewEVMAPIQueryHandler[K providertypes.ResponseKey, V providertypes.ResponseValue](
	logger *zap.Logger,
	cfg config.APIConfig,
	requestHandler EVMRequestHandler[K, V],
	metrics metrics.APIMetrics,
) (handlers.APIQueryHandler[K, V], error) {
	if err := cfg.ValidateBasic(); err != nil {
		return nil, fmt.Errorf("invalid provider config: %w", err)
	}

	if !cfg.Enabled {
		return nil, fmt.Errorf("api query handler is not enabled for the provider")
	}

	if logger == nil {
		return nil, fmt.Errorf("no logger specified for api query handler")
	}

	if requestHandler == nil {
		return nil, fmt.Errorf("no request handler specified for api query handler")
	}

	if metrics == nil {
		return nil, fmt.Errorf("no metrics specified for api query handler")
	}

	return &EVMAPIQueryHandlerImpl[K, V]{
		logger:         logger.With(zap.String("api_data_handler", cfg.Name)),
		config:         cfg,
		requestHandler: requestHandler,
		metrics:        metrics,
	}, nil
}

// Query is used to query the API data provider for the given IDs. This method blocks
// until all responses have been sent to the response channel. Query will only
// make N concurrent requests at a time, where N is the capacity of the response channel.
func (h *EVMAPIQueryHandlerImpl[K, V]) Query(
	ctx context.Context,
	ids []K,
	responseCh chan<- providertypes.GetResponse[K, V],
) {
	if len(ids) == 0 {
		h.logger.Debug("no ids to query")
		return
	}

	// Observe the total amount of time it takes to fulfill the request(s).
	h.logger.Debug("starting api query handler")
	start := time.Now().UTC()
	defer func() {
		if r := recover(); r != nil {
			h.logger.Error("panic in api query handler", zap.Any("panic", r))
		}

		h.metrics.ObserveProviderResponseLatency(h.config.Name, time.Since(start))
		h.logger.Debug("finished api query handler")
	}()

	// Set the concurrency limit based on the capacity of the channel. This is done
	// to ensure the query handler does not exceed the rate limit parameters of the
	// data provider.
	wg := errgroup.Group{}
	wg.SetLimit(cap(responseCh))
	h.logger.Debug("setting concurrency limit", zap.Int("limit", cap(responseCh)))

	// If our task is atomic, we can make a single request for all the IDs. Otherwise,
	// we need to make a request for each ID.
	var tasks []func() error
	if h.config.Atomic {
		tasks = append(tasks, h.subTask(ctx, ids, responseCh))
	} else {
		for i := 0; i < len(ids); i++ {
			id := ids[i]
			tasks = append(tasks, h.subTask(ctx, []K{id}, responseCh))
		}
	}

	// Block each task until the wait group has capacity to accept a new response.
	for _, task := range tasks {
		wg.Go(task)
	}

	// Wait for all tasks to complete.
	if err := wg.Wait(); err != nil {
		h.logger.Error("error querying ids", zap.Error(err))
	}
}

// subTask is the subtask that is used to query the data provider for the given IDs,
// parse the response, and write the response to the response channel.
func (h *EVMAPIQueryHandlerImpl[K, V]) subTask(
	ctx context.Context,
	ids []K,
	responseCh chan<- providertypes.GetResponse[K, V],
) func() error {
	return func() error {
		defer func() {
			// Recover from any panics that occur.
			if r := recover(); r != nil {
				h.logger.Error("panic occurred in subtask", zap.Any("panic", r), zap.Any("ids", ids))
			}

			h.logger.Debug("finished subtask", zap.Any("ids", ids))
		}()

		h.logger.Debug("starting subtask", zap.Any("ids", ids))

		// Make the request.
		h.writeResponse(responseCh, h.requestHandler.FetchPrices(ctx, h.config.URL, ids))
		return nil
	}
}

// writeResponse is used to write the response to the response channel.
func (h *EVMAPIQueryHandlerImpl[K, V]) writeResponse(
	responseCh chan<- providertypes.GetResponse[K, V],
	response providertypes.GetResponse[K, V],
) {
	responseCh <- response
	h.logger.Debug("wrote response", zap.String("response", response.String()))

	// Update the metrics.
	for id := range response.Resolved {
		h.metrics.AddProviderResponse(h.config.Name, strings.ToLower(id.String()), metrics.Success)
	}
	for id, err := range response.UnResolved {
		h.metrics.AddProviderResponse(h.config.Name, strings.ToLower(id.String()), metrics.StatusFromError(err))
	}
}
