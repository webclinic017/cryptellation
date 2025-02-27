package backtests

import (
	"context"
	"fmt"

	"github.com/digital-feather/cryptellation/services/backtests/internal/application/domains/backtest"
	"github.com/digital-feather/cryptellation/services/backtests/pkg/models/order"
	candlesticks "github.com/digital-feather/cryptellation/services/candlesticks/pkg/client"
)

func (b Backtests) CreateOrder(ctx context.Context, backtestId uint, order order.Order) error {
	return b.db.LockedBacktest(ctx, backtestId, func(bt *backtest.Backtest) error {
		list, err := b.csClient.ReadCandlesticks(ctx, candlesticks.ReadCandlestickPayload{
			ExchangeName: order.ExchangeName,
			PairSymbol:   order.PairSymbol,
			Period:       bt.PeriodBetweenEvents,
			Start:        bt.CurrentCsTick.Time,
			End:          bt.CurrentCsTick.Time,
			Limit:        0,
		})
		if err != nil {
			return fmt.Errorf("could not get candlesticks from service: %w", err)
		}

		tcs, notEmpty := list.First()
		if !notEmpty {
			return backtest.ErrNoDataForOrderValidation
		}

		if err := bt.AddOrder(order, tcs.Candlestick); err != nil {
			return err
		}

		return nil
	})
}
