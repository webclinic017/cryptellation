package binance

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/digital-feather/cryptellation/services/candlesticks/pkg/models/pairs"
	"github.com/digital-feather/cryptellation/services/ticks/internal/infrastructure/exchanges"
	"github.com/digital-feather/cryptellation/services/ticks/pkg/models/tick"

	client "github.com/adshao/go-binance/v2"
)

const Name = "binance"

type Service struct {
}

func New() (*Service, error) {
	return &Service{}, nil
}

func (s *Service) ListenSymbol(symbol string) (chan tick.Tick, chan struct{}, error) {
	binanceSymbol, err := toBinanceSymbol(symbol)
	if err != nil {
		return nil, nil, err
	}

	tickChan := make(chan tick.Tick)
	_, stop, err := client.WsBookTickerServe(binanceSymbol, func(event *client.WsBookTickerEvent) {
		ask, err := strconv.ParseFloat(event.BestAskPrice, 64)
		if err != nil {
			log.Println(err)
			return
		}

		bid, err := strconv.ParseFloat(event.BestBidPrice, 64)
		if err != nil {
			log.Println(err)
			return
		}

		t := tick.Tick{
			Time:       time.Now().UTC(),
			Exchange:   "binance",
			PairSymbol: symbol,
			Price:      float64(ask+bid) / 2,
		}

		tickChan <- t
	}, nil)

	// TODO: manage when error or done

	return tickChan, stop, err
}

func toBinanceSymbol(symbol string) (string, error) {
	base, quote, err := pairs.ParsePairSymbol(symbol)
	return fmt.Sprintf("%s%s", base, quote), err
}

func (s *Service) Name() string {
	return exchanges.BinanceName
}
