// Generate code for mock
//go:generate go run -mod=mod github.com/golang/mock/mockgen -source=adapter.go -destination=mock.gen.go -package db

package db

import (
	"context"
	"time"

	"github.com/digital-feather/cryptellation/services/backtests/internal/application/domains/backtest"
)

const (
	Expiration = 3 * time.Second
	RetryDelay = 100 * time.Millisecond
	Tries      = 20
)

type LockedBacktestCallback func(bt *backtest.Backtest) error

type Adapter interface {
	CreateBacktest(ctx context.Context, bt *backtest.Backtest) error
	ReadBacktest(ctx context.Context, id uint) (backtest.Backtest, error)
	UpdateBacktest(ctx context.Context, bt backtest.Backtest) error
	DeleteBacktest(ctx context.Context, bt backtest.Backtest) error

	LockedBacktest(ctx context.Context, id uint, fn LockedBacktestCallback) error
}
