package sql

import (
	"testing"

	"github.com/digital-feather/cryptellation/services/backtests/internal/application/ports/db/test"
	"github.com/stretchr/testify/suite"
)

func TestBacktestSuite(t *testing.T) {
	suite.Run(t, new(BacktestSuite))
}

type BacktestSuite struct {
	test.BacktestSuite
}

func (suite *BacktestSuite) SetupTest() {
	defer tmpEnvVar("SQLDB_DATABASE", "backtests")()

	db, err := New()
	suite.Require().NoError(err)
	suite.Require().NoError(db.Reset())

	suite.DB = db
}
