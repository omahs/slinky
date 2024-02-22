package handlers

import (
	"context"
	providertypes "github.com/skip-mev/slinky/providers/types"
)

type EVMRequestHandler[K providertypes.ResponseKey, V providertypes.ResponseValue] interface {
	FetchPrices(
		ctx context.Context,
		url string,
		ids []K,
	) providertypes.GetResponse[K, V]
}
