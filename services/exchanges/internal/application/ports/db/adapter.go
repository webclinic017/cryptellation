// Generate code for mock
//go:generate go run -mod=mod github.com/golang/mock/mockgen -source=adapter.go -destination=mock.gen.go -package db

package db

import (
	"context"

	"github.com/digital-feather/cryptellation/services/exchanges/pkg/models/exchange"
)

type Adapter interface {
	CreateExchanges(ctx context.Context, exchanges ...exchange.Exchange) error
	ReadExchanges(ctx context.Context, names ...string) ([]exchange.Exchange, error)
	UpdateExchanges(ctx context.Context, exchanges ...exchange.Exchange) error
	DeleteExchanges(ctx context.Context, names ...string) error
}
