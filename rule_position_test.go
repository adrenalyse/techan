package techan

import (
	"github.com/ericlagergren/decimal"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPositionNewRule(t *testing.T) {
	t.Run("returns true when position new", func(t *testing.T) {
		record := NewTradingRecord()
		rule := PositionNewRule{}

		assert.True(t, rule.IsSatisfied(0, record))
	})

	t.Run("returns false when position open", func(t *testing.T) {
		record := NewTradingRecord()

		record.Operate(Order{
			Side:   BUY,
			Amount: decimal.New(1, 0),
			Price:  decimal.New(1, 0),
		})

		rule := PositionNewRule{}

		assert.False(t, rule.IsSatisfied(0, record))
	})
}

func TestPositionOpenRule(t *testing.T) {
	t.Run("returns false when position new", func(t *testing.T) {
		record := NewTradingRecord()

		rule := PositionOpenRule{}

		assert.False(t, rule.IsSatisfied(0, record))
	})

	t.Run("returns true when position open", func(t *testing.T) {
		record := NewTradingRecord()

		record.Operate(Order{
			Side:   BUY,
			Amount: decimal.New(1, 0),
			Price:  decimal.New(1, 0),
		})

		rule := PositionOpenRule{}

		assert.True(t, rule.IsSatisfied(0, record))
	})
}
