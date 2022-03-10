package application

import (
	"github.com/cryptellation/cryptellation/services/exchanges/internal/application/commands"
)

type Application struct {
	Commands Commands
	Queries  Queries
}

type Commands struct {
	CachedReadExchanges commands.CachedReadExchangesHandler
}

type Queries struct {
}
