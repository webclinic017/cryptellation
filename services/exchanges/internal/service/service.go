package service

import (
	"github.com/digital-feather/cryptellation/services/exchanges/internal/adapters/db/cockroach"
	"github.com/digital-feather/cryptellation/services/exchanges/internal/adapters/exchanges"
	"github.com/digital-feather/cryptellation/services/exchanges/internal/adapters/exchanges/binance"
	app "github.com/digital-feather/cryptellation/services/exchanges/internal/application"
	"github.com/digital-feather/cryptellation/services/exchanges/internal/application/commands"
)

func NewApplication() (app.Application, error) {
	binanceService, err := binance.New()
	if err != nil {
		return app.Application{}, err
	}

	services := map[string]exchanges.Port{
		exchanges.Binance.Name: binanceService,
	}

	return newApplication(services)
}

func newMockApplication() (app.Application, error) {
	services := map[string]exchanges.Port{
		"mock_exchange": MockExchangeService{},
	}

	return newApplication(services)
}

func newApplication(services map[string]exchanges.Port) (app.Application, error) {
	repository, err := cockroach.New()
	if err != nil {
		return app.Application{}, err
	}

	return app.Application{
		Commands: app.Commands{
			CachedReadExchanges: commands.NewCachedReadExchangesHandler(repository, services),
		},
	}, nil
}
