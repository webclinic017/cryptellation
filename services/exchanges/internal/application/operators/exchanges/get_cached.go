package exchanges

import (
	"context"
	"fmt"
	"time"

	domain "github.com/digital-feather/cryptellation/services/exchanges/internal/application/domain/exchange"
	"github.com/digital-feather/cryptellation/services/exchanges/pkg/models/exchange"
	"golang.org/x/xerrors"
)

func (e Exchanges) GetCached(ctx context.Context, expiration *time.Duration, names ...string) ([]exchange.Exchange, error) {
	dbExchanges, err := e.db.ReadExchanges(ctx, names...)
	if err != nil {
		return nil, fmt.Errorf("handling exchanges from db reading: %w", err)
	}

	toSync, err := domain.GetExpiredExchangesNames(names, dbExchanges, expiration)
	if err != nil {
		return nil, fmt.Errorf("determining exchanges to synchronize: %w", err)
	}

	synced, err := e.getExchangeFromServices(ctx, toSync...)
	if err != nil {
		return nil, err
	}

	err = e.upsertExchanges(ctx, dbExchanges, synced)
	if err != nil {
		return nil, err
	}

	mappedExchanges := exchange.ArrayToMap(dbExchanges)
	for _, exch := range synced {
		mappedExchanges[exch.Name] = exch
	}

	return exchange.MapToArray(mappedExchanges), nil
}

func (e Exchanges) getExchangeFromServices(ctx context.Context, toSync ...string) ([]exchange.Exchange, error) {
	synced := make([]exchange.Exchange, 0, len(toSync))
	for _, name := range toSync {
		service, ok := e.services[name]
		if !ok {
			return nil, xerrors.New(fmt.Sprintf("inexistant exchange service %q", name))
		}

		exch, err := service.Infos(ctx)
		if err != nil {
			return nil, err
		}

		synced = append(synced, exch)
	}

	return synced, nil
}

func (e Exchanges) upsertExchanges(ctx context.Context, dbExchanges, toUpsert []exchange.Exchange) error {
	toCreate := make([]exchange.Exchange, 0, len(toUpsert))
	toUpdate := make([]exchange.Exchange, 0, len(toUpsert))
	mappedDbExchanges := exchange.ArrayToMap(dbExchanges)
	for _, exch := range toUpsert {
		if _, ok := mappedDbExchanges[exch.Name]; ok {
			toUpdate = append(toUpdate, exch)
		} else {
			toCreate = append(toCreate, exch)
		}
	}

	if len(toCreate) > 0 {
		if err := e.db.CreateExchanges(ctx, toCreate...); err != nil {
			return err
		}
	}

	if len(toUpdate) > 0 {
		if err := e.db.UpdateExchanges(ctx, toUpdate...); err != nil {
			return err
		}
	}

	return nil
}
