package tests

import (
	"context"
	"time"

	"github.com/digital-feather/cryptellation/services/candlesticks/internal/application/ports/db"
	"github.com/digital-feather/cryptellation/services/candlesticks/pkg/models/candlestick"
	"github.com/digital-feather/cryptellation/services/candlesticks/pkg/models/period"
	"github.com/stretchr/testify/suite"
)

type CandlesticksSuite struct {
	suite.Suite
	DB db.Adapter
}

func (suite *CandlesticksSuite) TestCreate() {
	list := candlestick.NewEmptyList(candlestick.ListID{
		ExchangeName: "exchange",
		PairSymbol:   "ETH-USDC",
		Period:       period.M1,
	})
	t, err := time.Parse(time.RFC3339, "1993-11-15T11:29:00Z")
	suite.Require().NoError(err)
	suite.Require().NoError(list.Set(t, candlestick.Candlestick{
		Open:       1,
		Low:        0.5,
		High:       2,
		Close:      1.5,
		Volume:     1000,
		Uncomplete: true,
	}))
	recvList := candlestick.NewEmptyList(candlestick.ListID{
		ExchangeName: "exchange",
		PairSymbol:   "ETH-USDC",
		Period:       period.M1,
	})

	suite.Require().NoError(suite.DB.CreateCandlesticks(context.Background(), list))
	err = suite.DB.ReadCandlesticks(context.Background(), recvList, t.Add(-time.Hour), t.Add(time.Hour), 0)
	suite.Require().NoError(err)

	suite.Require().Equal(1, recvList.Len())
	cs, _ := list.Get(t)
	rcs, exists := recvList.Get(t)
	suite.True(exists)
	suite.True(cs == rcs)
}

func (suite *CandlesticksSuite) TestCreateTwice() {
	list := candlestick.NewEmptyList(candlestick.ListID{
		ExchangeName: "exchange",
		PairSymbol:   "ETH-USDC",
		Period:       period.M1,
	})
	t, err := time.Parse(time.RFC3339, "1993-11-15T11:29:00Z")
	suite.Require().NoError(err)
	suite.Require().NoError(list.Set(t, candlestick.Candlestick{
		Open:   1,
		Low:    0.5,
		High:   2,
		Close:  1.5,
		Volume: 1000,
	}))

	suite.Require().NoError(suite.DB.CreateCandlesticks(context.Background(), list))
	suite.Require().Error(suite.DB.CreateCandlesticks(context.Background(), list))
}

func (suite *CandlesticksSuite) TestRead() {
	// Create targeted exchange, pair, period
	list := candlestick.NewEmptyList(candlestick.ListID{
		ExchangeName: "exchange",
		PairSymbol:   "ETH-USDC",
		Period:       period.M1,
	})

	t1, err := time.Parse(time.RFC3339, "1993-11-15T11:29:00Z")
	suite.Require().NoError(err)
	suite.Require().NoError(list.Set(t1, candlestick.Candlestick{
		Open:   1,
		Low:    0.5,
		High:   2,
		Close:  1.5,
		Volume: 1000,
	}))

	t2, err := time.Parse(time.RFC3339, "1993-11-15T11:30:00Z")
	suite.Require().NoError(err)
	suite.Require().NoError(list.Set(t2, candlestick.Candlestick{
		Open:   2,
		Low:    1,
		High:   4,
		Close:  3,
		Volume: 2000,
	}))

	t3, err := time.Parse(time.RFC3339, "1993-11-15T11:31:00Z")
	suite.Require().NoError(err)
	suite.Require().NoError(list.Set(t3, candlestick.Candlestick{
		Open:   3,
		Low:    1.5,
		High:   6,
		Close:  4.5,
		Volume: 3000,
	}))

	t4, err := time.Parse(time.RFC3339, "1993-11-15T11:32:00Z")
	suite.Require().NoError(err)
	suite.Require().NoError(list.Set(t4, candlestick.Candlestick{
		Open:   4,
		Low:    2,
		High:   8,
		Close:  6,
		Volume: 4000,
	}))

	suite.Require().NoError(suite.DB.CreateCandlesticks(context.Background(), list))

	// Create other exchange
	otherExchangeList := candlestick.NewEmptyList(candlestick.ListID{
		ExchangeName: "exchange2",
		PairSymbol:   "ETH-USDC",
		Period:       period.M1,
	})
	suite.Require().NoError(err)
	suite.Require().NoError(otherExchangeList.Set(t2, candlestick.Candlestick{
		Open:   1,
		Low:    1,
		High:   1,
		Close:  1,
		Volume: 1,
	}))
	suite.Require().NoError(suite.DB.CreateCandlesticks(context.Background(), otherExchangeList))

	// Create other pair
	otherPairList := candlestick.NewEmptyList(candlestick.ListID{
		ExchangeName: "exchange",
		PairSymbol:   "BTC-USDC",
		Period:       period.M1,
	})
	suite.Require().NoError(err)
	suite.Require().NoError(otherPairList.Set(t2, candlestick.Candlestick{
		Open:   2,
		Low:    2,
		High:   2,
		Close:  2,
		Volume: 2,
	}))
	suite.Require().NoError(suite.DB.CreateCandlesticks(context.Background(), otherPairList))

	// Create other period
	otherPeriodList := candlestick.NewEmptyList(candlestick.ListID{
		ExchangeName: "exchange",
		PairSymbol:   "ETH-USDC",
		Period:       period.M15,
	})
	suite.Require().NoError(err)
	suite.Require().NoError(otherPeriodList.Set(t2, candlestick.Candlestick{
		Open:   3,
		Low:    3,
		High:   3,
		Close:  3,
		Volume: 3,
	}))
	suite.Require().NoError(suite.DB.CreateCandlesticks(context.Background(), otherPeriodList))

	// Read only the two centered candles
	recvList := candlestick.NewEmptyList(candlestick.ListID{
		ExchangeName: "exchange",
		PairSymbol:   "ETH-USDC",
		Period:       period.M1,
	})
	suite.Require().NoError(suite.DB.ReadCandlesticks(context.Background(), recvList, t2, t3, 0))
	suite.Require().Equal(2, recvList.Len())
	c, _ := list.Get(t2)
	rc, exists := recvList.Get(t2)
	suite.True(exists)
	suite.Require().Equal(c, rc)

	c, _ = list.Get(t3)
	rc, exists = recvList.Get(t3)
	suite.True(exists)
	suite.Require().Equal(c, rc)

	// Check others
	recvList = candlestick.NewEmptyList(candlestick.ListID{
		ExchangeName: "exchange2",
		PairSymbol:   "ETH-USDC",
		Period:       period.M1,
	})
	suite.Require().NoError(suite.DB.ReadCandlesticks(context.Background(), recvList, t2, t3, 0))
	c, _ = otherExchangeList.Get(t2)
	rc, exists = recvList.Get(t2)
	suite.True(exists)
	suite.Require().Equal(c, rc)

	recvList = candlestick.NewEmptyList(candlestick.ListID{
		ExchangeName: "exchange",
		PairSymbol:   "BTC-USDC",
		Period:       period.M1,
	})
	suite.Require().NoError(suite.DB.ReadCandlesticks(context.Background(), recvList, t2, t3, 0))
	c, _ = otherPairList.Get(t2)
	rc, exists = recvList.Get(t2)
	suite.True(exists)
	suite.Require().Equal(c, rc)

	recvList = candlestick.NewEmptyList(candlestick.ListID{
		ExchangeName: "exchange",
		PairSymbol:   "ETH-USDC",
		Period:       period.M15,
	})
	suite.Require().NoError(suite.DB.ReadCandlesticks(context.Background(), recvList, t2, t3, 0))
	c, _ = otherPeriodList.Get(t2)
	rc, exists = recvList.Get(t2)
	suite.True(exists)
	suite.Require().Equal(c, rc)
}

func (suite *CandlesticksSuite) TestReadLimit() {
	// Create targeted exchange, pair, period
	list := candlestick.NewEmptyList(candlestick.ListID{
		ExchangeName: "exchange",
		PairSymbol:   "ETH-USDC",
		Period:       period.M1,
	})

	t1, err := time.Parse(time.RFC3339, "1993-11-15T11:29:00Z")
	suite.Require().NoError(err)
	suite.Require().NoError(list.Set(t1, candlestick.Candlestick{
		Open:   1,
		Low:    0.5,
		High:   2,
		Close:  1.5,
		Volume: 1000,
	}))

	t2, err := time.Parse(time.RFC3339, "1993-11-15T11:30:00Z")
	suite.Require().NoError(err)
	suite.Require().NoError(list.Set(t2, candlestick.Candlestick{
		Open:   2,
		Low:    1,
		High:   4,
		Close:  3,
		Volume: 2000,
	}))

	t3, err := time.Parse(time.RFC3339, "1993-11-15T11:31:00Z")
	suite.Require().NoError(err)
	suite.Require().NoError(list.Set(t3, candlestick.Candlestick{
		Open:   3,
		Low:    1.5,
		High:   6,
		Close:  4.5,
		Volume: 3000,
	}))

	t4, err := time.Parse(time.RFC3339, "1993-11-15T11:32:00Z")
	suite.Require().NoError(err)
	suite.Require().NoError(list.Set(t4, candlestick.Candlestick{
		Open:   4,
		Low:    2,
		High:   8,
		Close:  6,
		Volume: 4000,
	}))

	suite.Require().NoError(suite.DB.CreateCandlesticks(context.Background(), list))

	// Read only the 2 first
	recvList := candlestick.NewEmptyList(candlestick.ListID{
		ExchangeName: "exchange",
		PairSymbol:   "ETH-USDC",
		Period:       period.M1,
	})
	suite.Require().NoError(suite.DB.ReadCandlesticks(context.Background(), recvList, t1, t4, 2))
	suite.Require().Equal(2, recvList.Len())
	c, _ := list.Get(t1)
	rc, exists := recvList.Get(t1)
	suite.True(exists)
	suite.Require().Equal(c, rc)

	c, _ = list.Get(t2)
	rc, exists = recvList.Get(t2)
	suite.True(exists)
	suite.Require().Equal(c, rc)
}

func (suite *CandlesticksSuite) TestReadEmpty() {
	t, err := time.Parse(time.RFC3339, "1993-11-15T11:29:00Z")
	suite.Require().NoError(err)
	recvList := candlestick.NewEmptyList(candlestick.ListID{
		ExchangeName: "exchange",
		PairSymbol:   "ETH-USDC",
		Period:       period.M1,
	})
	suite.Require().NoError(suite.DB.ReadCandlesticks(context.Background(), recvList, t.Add(-time.Hour), t.Add(time.Hour), 0))
	suite.Require().Equal(0, recvList.Len())
}

func (suite *CandlesticksSuite) TestUpdate() {
	list := candlestick.NewEmptyList(candlestick.ListID{
		ExchangeName: "exchange",
		PairSymbol:   "ETH-USDC",
		Period:       period.M1,
	})
	t, err := time.Parse(time.RFC3339, "1993-11-15T11:29:00Z")
	suite.Require().NoError(err)
	suite.Require().NoError(list.Set(t, candlestick.Candlestick{
		Open:       1,
		Low:        0.5,
		High:       2,
		Close:      1.5,
		Volume:     1000,
		Uncomplete: true,
	}))
	suite.Require().NoError(suite.DB.CreateCandlesticks(context.Background(), list))

	update := candlestick.NewEmptyList(candlestick.ListID{
		ExchangeName: "exchange",
		PairSymbol:   "ETH-USDC",
		Period:       period.M1,
	})
	suite.Require().NoError(update.Set(t, candlestick.Candlestick{
		Open:       2,
		Low:        1,
		High:       4,
		Close:      3,
		Volume:     2000,
		Uncomplete: false,
	}))
	suite.Require().NoError(suite.DB.UpdateCandlesticks(context.Background(), update))
	receivedList := candlestick.NewEmptyList(candlestick.ListID{
		ExchangeName: "exchange",
		PairSymbol:   "ETH-USDC",
		Period:       period.M1,
	})
	err = suite.DB.ReadCandlesticks(context.Background(), receivedList, t.Add(-time.Hour), t.Add(time.Hour), 0)
	suite.Require().NoError(err)

	suite.Require().Equal(1, receivedList.Len())
	cs, _ := update.Get(t)
	rcs, exists := receivedList.Get(t)
	suite.True(exists)
	suite.Require().Equal(cs, rcs)
}

func (suite *CandlesticksSuite) TestUpdateInexistantTwice() {
	list := candlestick.NewEmptyList(candlestick.ListID{
		ExchangeName: "exchange",
		PairSymbol:   "ETH-USDC",
		Period:       period.M1,
	})
	t, err := time.Parse(time.RFC3339, "1993-11-15T11:29:00Z")
	suite.Require().NoError(err)
	suite.Require().NoError(list.Set(t, candlestick.Candlestick{
		Open:   1,
		Low:    0.5,
		High:   2,
		Close:  1.5,
		Volume: 1000,
	}))

	suite.Require().Error(suite.DB.UpdateCandlesticks(context.Background(), list))
	suite.Require().Error(suite.DB.UpdateCandlesticks(context.Background(), list))
}

func (suite *CandlesticksSuite) TestDelete() {
	list := candlestick.NewEmptyList(candlestick.ListID{
		ExchangeName: "exchange",
		PairSymbol:   "ETH-USDC",
		Period:       period.M1,
	})
	for i := 0; i < 10; i++ {
		suite.Require().NoError(list.Set(time.Unix(int64(i*int(period.M1.Duration().Seconds())), 0), candlestick.Candlestick{
			Open:   1 + float64(i),
			Low:    0.5 + float64(i),
			High:   2 + float64(i),
			Close:  1.5 + float64(i),
			Volume: 1000 * float64(i),
		}))
	}
	suite.Require().NoError(suite.DB.CreateCandlesticks(context.Background(), list))

	// Remove half the data
	delete := candlestick.NewEmptyList(candlestick.ListID{
		ExchangeName: "exchange",
		PairSymbol:   "ETH-USDC",
		Period:       period.M1,
	})
	suite.Require().NoError(delete.Set(time.Unix(int64(0*period.M1.Duration().Seconds()), 0), candlestick.Candlestick{}))
	suite.Require().NoError(delete.Set(time.Unix(int64(1*period.M1.Duration().Seconds()), 0), candlestick.Candlestick{}))
	suite.Require().NoError(delete.Set(time.Unix(int64(2*period.M1.Duration().Seconds()), 0), candlestick.Candlestick{}))
	suite.Require().NoError(delete.Set(time.Unix(int64(3*period.M1.Duration().Seconds()), 0), candlestick.Candlestick{}))
	suite.Require().NoError(delete.Set(time.Unix(int64(4*period.M1.Duration().Seconds()), 0), candlestick.Candlestick{}))
	suite.Require().NoError(suite.DB.DeleteCandlesticks(context.Background(), delete))

	// Check staying data
	tEnd := time.Unix(10*int64(period.M1.Duration().Seconds()), 0)
	tStart := time.Unix(5*int64(period.M1.Duration().Seconds()), 0)
	receivedList := candlestick.NewEmptyList(candlestick.ListID{
		ExchangeName: "exchange",
		PairSymbol:   "ETH-USDC",
		Period:       period.M1,
	})
	err := suite.DB.ReadCandlesticks(context.Background(), receivedList, tStart, tEnd, 0)
	suite.Require().NoError(err)
	suite.Require().Equal(list.Len()-5, receivedList.Len())
	suite.Require().NoError(receivedList.Loop(func(ts time.Time, cs candlestick.Candlestick) (bool, error) {
		origCS, exists := list.Get(ts)
		suite.True(exists)
		suite.Require().Equal(origCS, cs)
		return false, nil
	}))
}
