package ticks

import "context"

func (t Ticks) Unregister(ctx context.Context, exchange, pairSymbol string) (int64, error) {
	count, err := t.db.DecrementSymbolListenerSubscribers(ctx, exchange, pairSymbol)
	if err != nil {
		return count, err
	}

	return count, nil
}
