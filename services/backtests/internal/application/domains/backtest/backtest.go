package backtest

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/digital-feather/cryptellation/services/backtests/pkg/models/account"
	"github.com/digital-feather/cryptellation/services/backtests/pkg/models/event"
	"github.com/digital-feather/cryptellation/services/backtests/pkg/models/order"
	"github.com/digital-feather/cryptellation/services/candlesticks/pkg/models/candlestick"
	"github.com/digital-feather/cryptellation/services/candlesticks/pkg/models/pairs"
	"github.com/digital-feather/cryptellation/services/candlesticks/pkg/models/period"
)

var (
	ErrTickSubscriptionAlreadyExists = errors.New("tick-subscription-already-exists")
	ErrInvalidExchange               = errors.New("invalid-exchange")
	ErrNotEnoughAsset                = errors.New("no-enough-asset")
	ErrNoDataForOrderValidation      = errors.New("no-data-for-order-validation")
	ErrStartAfterEnd                 = errors.New("start-after-end")
)

// Current tick based on candlestick step
type CurrentCsTick struct {
	Time      time.Time
	PriceType candlestick.PriceType
}

type Backtest struct {
	ID                  uint
	StartTime           time.Time
	CurrentCsTick       CurrentCsTick
	EndTime             time.Time
	Accounts            map[string]account.Account
	PeriodBetweenEvents period.Symbol
	TickSubscriptions   []event.TickSubscription
	Orders              []order.Order
}

type NewPayload struct {
	Accounts              map[string]account.Account
	StartTime             time.Time
	EndTime               *time.Time
	DurationBetweenEvents *time.Duration
}

func (payload *NewPayload) EmptyFieldsToDefault() *NewPayload {
	if payload.EndTime == nil {
		payload.EndTime = defaultEndTime()
	}

	if payload.DurationBetweenEvents == nil {
		d := time.Minute
		payload.DurationBetweenEvents = &d
	}

	return payload
}

func (payload NewPayload) Validate() error {
	if !payload.StartTime.Before(*payload.EndTime) {
		return ErrStartAfterEnd
	}

	for exchangeName, a := range payload.Accounts {
		if exchangeName == "" {
			return fmt.Errorf("error with exchange %q in new backtest payload: %w", exchangeName, ErrInvalidExchange)
		}

		if err := a.Validate(); err != nil {
			return err
		}
	}

	return nil
}

func defaultEndTime() *time.Time {
	t := time.Now()
	return &t
}

func New(ctx context.Context, payload NewPayload) (Backtest, error) {
	if err := payload.EmptyFieldsToDefault().Validate(); err != nil {
		return Backtest{}, err
	}

	per, err := period.FromDuration(*payload.DurationBetweenEvents)
	if err != nil {
		return Backtest{}, fmt.Errorf("invalid duration between events: %w", err)
	}

	return Backtest{
		StartTime: payload.StartTime,
		CurrentCsTick: CurrentCsTick{
			Time:      payload.StartTime,
			PriceType: candlestick.PriceTypeIsOpen,
		},
		EndTime:             *payload.EndTime,
		Accounts:            payload.Accounts,
		PeriodBetweenEvents: per,
		TickSubscriptions:   make([]event.TickSubscription, 0),
		Orders:              make([]order.Order, 0),
	}, nil
}

func (bt Backtest) MarshalBinary() ([]byte, error) {
	return json.Marshal(bt)
}

func (bt *Backtest) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, bt)
}

func (bt *Backtest) Advance() (done bool) {
	return bt.advanceThroughTicks()
}

func (bt *Backtest) advanceThroughTicks() (done bool) {
	switch bt.CurrentCsTick.PriceType {
	case candlestick.PriceTypeIsOpen:
		bt.CurrentCsTick.PriceType = candlestick.PriceTypeIsHigh
	case candlestick.PriceTypeIsHigh:
		bt.CurrentCsTick.PriceType = candlestick.PriceTypeIsLow
	case candlestick.PriceTypeIsLow:
		bt.CurrentCsTick.PriceType = candlestick.PriceTypeIsClose
	case candlestick.PriceTypeIsClose:
		bt.SetCurrentTime(bt.CurrentCsTick.Time.Add(bt.PeriodBetweenEvents.Duration()))
	default:
		bt.CurrentCsTick.PriceType = candlestick.PriceTypeIsOpen
	}

	return bt.Done()
}

func (bt Backtest) Done() bool {
	return !bt.CurrentCsTick.Time.Before(bt.EndTime)
}

func (bt *Backtest) SetCurrentTime(ts time.Time) {
	bt.CurrentCsTick.Time = ts
	bt.CurrentCsTick.PriceType = candlestick.PriceTypeIsOpen
}

func (bt *Backtest) CreateTickSubscription(exchangeName string, pairSymbol string) (event.TickSubscription, error) {
	for _, ts := range bt.TickSubscriptions {
		if ts.ExchangeName == exchangeName && ts.PairSymbol == pairSymbol {
			return event.TickSubscription{}, ErrTickSubscriptionAlreadyExists
		}
	}

	s := event.TickSubscription{
		ID:           len(bt.TickSubscriptions),
		ExchangeName: exchangeName,
		PairSymbol:   pairSymbol,
	}
	bt.TickSubscriptions = append(bt.TickSubscriptions, s)

	return s, nil
}

func (bt *Backtest) AddOrder(ord order.Order, cs candlestick.Candlestick) error {
	exchangeAccount, ok := bt.Accounts[ord.ExchangeName]
	if !ok {
		return fmt.Errorf("error with orders exchange %q: %w", ord.ExchangeName, ErrInvalidExchange)
	}

	baseSymbol, quoteSymbol, err := pairs.ParsePairSymbol(ord.PairSymbol)
	if err != nil {
		return fmt.Errorf("error when parsing order pair symbol: %w", err)
	}

	price := cs.PriceByType(bt.CurrentCsTick.PriceType)
	quoteEquivalentQty := price * ord.Quantity
	if ord.Side == order.SideIsBuy {
		available, ok := exchangeAccount.Balances[quoteSymbol]
		if !ok {
			return ErrNotEnoughAsset
		} else if quoteEquivalentQty > available {
			return ErrNotEnoughAsset
		}

		bt.Accounts[ord.ExchangeName].Balances[quoteSymbol] -= quoteEquivalentQty
		bt.Accounts[ord.ExchangeName].Balances[baseSymbol] += ord.Quantity
	} else {
		available, ok := exchangeAccount.Balances[baseSymbol]
		if !ok {
			return ErrNotEnoughAsset
		} else if ord.Quantity > available {
			return ErrNotEnoughAsset
		}

		bt.Accounts[ord.ExchangeName].Balances[quoteSymbol] += quoteEquivalentQty
		bt.Accounts[ord.ExchangeName].Balances[baseSymbol] -= ord.Quantity
	}

	ord.ID = uint64(len(bt.Orders))
	ord.ExecutionTime = &bt.CurrentCsTick.Time
	ord.Price = price

	bt.Orders = append(bt.Orders, ord)
	return nil
}
