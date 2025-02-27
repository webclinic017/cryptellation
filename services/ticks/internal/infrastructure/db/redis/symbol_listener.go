package redis

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/go-redsync/redsync/v4"
	"github.com/go-redsync/redsync/v4/redis/goredis/v8"
)

const (
	redisKeySymbolListenerPrefix = "symbol-listeners-count"
	redisKeySymbolListener       = redisKeySymbolListenerPrefix + "-%s-%s"
)

type DB struct {
	client     *redis.Client
	lockClient *redsync.Redsync
}

func New() (*DB, error) {
	var c Config
	if err := c.Load().Validate(); err != nil {
		return nil, fmt.Errorf("loading redis config: %w", err)
	}

	client := redis.NewClient(&redis.Options{
		Addr:     c.Address,
		Password: c.Password, // no password set
		DB:       0,          // use default DB
	})

	// TODO Check connection

	pool := goredis.NewPool(client)
	ls := redsync.New(pool)

	return &DB{
		client:     client,
		lockClient: ls,
	}, nil
}

func (db *DB) IncrementSymbolListenerSubscribers(ctx context.Context, exchange, pairSymbol string) (int64, error) {
	key := fmt.Sprintf(redisKeySymbolListener, exchange, pairSymbol)
	return db.client.Incr(ctx, key).Result()
}

func (db *DB) DecrementSymbolListenerSubscribers(ctx context.Context, exchange, pairSymbol string) (int64, error) {
	key := fmt.Sprintf(redisKeySymbolListener, exchange, pairSymbol)
	return db.client.Decr(ctx, key).Result()
}

func (db *DB) GetSymbolListenerSubscribers(ctx context.Context, exchange, pairSymbol string) (int64, error) {
	key := fmt.Sprintf(redisKeySymbolListener, exchange, pairSymbol)
	content, err := db.client.Get(ctx, key).Result()
	if err != nil {
		return 0, err
	}

	return strconv.ParseInt(content, 10, 64)
}

func (db *DB) ClearAllSymbolListenersCount(ctx context.Context) error {
	keys, err := db.client.Keys(ctx, redisKeySymbolListenerPrefix+"*").Result()
	if err != nil {
		return err
	}

	for _, k := range keys {
		_, err := db.client.Set(ctx, k, 0, 0).Result()
		if err != nil {
			return err
		}
	}

	return nil
}

func (db *DB) ClearSymbolListenerSubscribers(ctx context.Context, exchange, pairSymbol string) error {
	key := fmt.Sprintf(redisKeySymbolListener, exchange, pairSymbol)
	_, err := db.client.Set(ctx, key, 0, time.Second).Result()
	return err
}
