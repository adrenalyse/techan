package techan

import (
	"github.com/ericlagergren/decimal"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewTradingRecord(t *testing.T) {
	record := NewTradingRecord()

	assert.Len(t, record.Trades, 0)
	assert.True(t, record.CurrentPosition().IsNew())
}

func TestTradingRecord_CurrentTrade(t *testing.T) {
	record := NewTradingRecord()

	yesterday := time.Now().Add(-time.Hour * 24)
	record.Operate(Order{
		Side:          BUY,
		Amount:        decimal.New(1, 0),
		Price:         decimal.New(2, 0),
		ExecutionTime: yesterday,
	})

	assert.EqualValues(t, "1", record.CurrentPosition().EntranceOrder().Amount.String())
	assert.EqualValues(t, "2", record.CurrentPosition().EntranceOrder().Price.String())
	assert.EqualValues(t, yesterday.UnixNano(),
		record.CurrentPosition().EntranceOrder().ExecutionTime.UnixNano())

	now := time.Now()
	record.Operate(Order{
		Side:          SELL,
		Amount:        decimal.New(3, 0),
		Price:         decimal.New(4, 0),
		ExecutionTime: now,
	})
	assert.True(t, record.CurrentPosition().IsNew())

	lastTrade := record.LastTrade()

	assert.EqualValues(t, "3", lastTrade.ExitOrder().Amount.String())
	assert.EqualValues(t, "4", lastTrade.ExitOrder().Price.String())
	assert.EqualValues(t, now.UnixNano(),
		lastTrade.ExitOrder().ExecutionTime.UnixNano())
}

func TestTradingRecord_Enter(t *testing.T) {
	t.Run("Does not add trades older than last trade", func(t *testing.T) {
		record := NewTradingRecord()

		now := time.Now()

		record.Operate(Order{
			Side:          BUY,
			Amount:        decimal.New(1, 0),
			Price:         decimal.New(2, 0),
			ExecutionTime: now,
		})

		record.Operate(Order{
			Side:          SELL,
			Amount:        decimal.New(2, 0),
			Price:         decimal.New(2, 0),
			ExecutionTime: now.Add(time.Minute),
		})

		record.Operate(Order{
			Side:          BUY,
			Amount:        decimal.New(2, 0),
			Price:         decimal.New(2, 0),
			ExecutionTime: now.Add(-time.Minute),
		})

		assert.True(t, record.CurrentPosition().IsNew())
		assert.Len(t, record.Trades, 1)
	})
}

func TestTradingRecord_Exit(t *testing.T) {
	t.Run("Does not add trades older than last trade", func(t *testing.T) {
		record := NewTradingRecord()

		now := time.Now()
		record.Operate(Order{

			Side:          BUY,
			Amount:        decimal.New(1, 0),
			Price:         decimal.New(2, 0),
			ExecutionTime: now,
		})

		record.Operate(Order{
			Side:          SELL,
			Amount:        decimal.New(2, 0),
			Price:         decimal.New(2, 0),
			ExecutionTime: now.Add(-time.Minute),
		})

		assert.True(t, record.CurrentPosition().IsOpen())
	})
}
