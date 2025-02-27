package sql

import (
	"context"
	"fmt"
	"time"

	"github.com/digital-feather/cryptellation/services/candlesticks/internal/application/ports/db"
	"github.com/digital-feather/cryptellation/services/candlesticks/internal/infrastructure/db/sql/entities"
	"github.com/digital-feather/cryptellation/services/candlesticks/pkg/models/candlestick"
)

func (d *DB) CreateCandlesticks(ctx context.Context, cs *candlestick.List) error {
	listCE := entities.FromModelListToEntityList(cs)
	return d.client.WithContext(ctx).Create(&listCE).Error
}

func (d *DB) ReadCandlesticks(ctx context.Context, cs *candlestick.List, start, end time.Time, limit uint) error {
	tx := d.client.Where(`
		exchange_name = ? AND
		pair_symbol = ? AND
		period_symbol = ? AND
		time BETWEEN ? AND ?`,
		cs.ExchangeName(),
		cs.PairSymbol(),
		cs.Period().String(),
		start, end)

	if limit != 0 {
		tx = tx.Limit(int(limit))
	}

	cse := []entities.Candlestick{}
	if err := tx.WithContext(ctx).Find(&cse).Error; err != nil {
		return err
	}

	for _, ce := range cse {
		_, _, _, t, m := ce.ToModel()
		if err := cs.Set(t, m); err != nil {
			return err
		}
	}

	return nil
}

func (d *DB) UpdateCandlesticks(ctx context.Context, cs *candlestick.List) error {
	listCE := entities.FromModelListToEntityList(cs)
	for _, ce := range listCE {
		tx := d.client.WithContext(ctx).
			Select("*").
			Model(&entities.Candlestick{}).
			Where("exchange_name = ?", ce.ExchangeName).
			Where("pair_symbol = ?", ce.PairSymbol).
			Where("period_symbol = ?", ce.PeriodSymbol).
			Where("time = ?", ce.Time).
			Updates(ce)

		if tx.Error != nil {
			return fmt.Errorf("updating candlestick %q: %w", ce.Time, tx.Error)
		} else if tx.RowsAffected == 0 {
			return db.ErrNotFound
		}
	}

	return nil
}

func (d *DB) DeleteCandlesticks(ctx context.Context, cs *candlestick.List) error {
	listCE := entities.FromModelListToEntityList(cs)
	return d.client.WithContext(ctx).Delete(&listCE).Error
}
