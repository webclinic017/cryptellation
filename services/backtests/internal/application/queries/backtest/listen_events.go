package queriesBacktest

import (
	"context"

	"github.com/cryptellation/cryptellation/services/backtests/internal/adapters/pubsub"
	"github.com/cryptellation/cryptellation/services/backtests/internal/domain/event"
)

type ListenEventsHandler struct {
	pubsub pubsub.Port
}

func NewListenEventsHandler(ps pubsub.Port) ListenEventsHandler {
	if ps == nil {
		panic("nil pubsub")
	}

	return ListenEventsHandler{
		pubsub: ps,
	}
}

func (h ListenEventsHandler) Handle(ctx context.Context, backtestId uint64) (<-chan event.Interface, error) {
	sub, err := h.pubsub.Subscribe(ctx, uint(backtestId))
	if err != nil {
		return nil, err
	}

	return sub.Channel(), nil
}
