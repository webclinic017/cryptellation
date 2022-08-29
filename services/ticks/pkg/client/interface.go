package client

import (
	"context"

	"github.com/digital-feather/cryptellation/services/ticks/pkg/models/tick"
)

type Client interface {
	Register(ctx context.Context, exchange, symbol string) error
	Unregister(ctx context.Context, exchange, symbol string) error
	Listen(symbol string) (<-chan tick.Tick, error)
}
