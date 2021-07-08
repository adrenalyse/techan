package techan

import (
	"github.com/ericlagergren/decimal"
	"testing"
	"time"

	"fmt"
	"strings"

	"github.com/stretchr/testify/assert"
)

func TestCandle_AddTrade(t *testing.T) {
	now := time.Now()
	candle := NewCandle(TimePeriod{
		Start: now,
		End:   now.Add(time.Minute),
	})

	candle.AddTrade(*decimal.New(1, 0), *decimal.New(2, 0)) // Open
	candle.AddTrade(*decimal.New(1, 0), *decimal.New(5, 0)) // High
	candle.AddTrade(*decimal.New(1, 0), *decimal.New(1, 0)) // Low
	candle.AddTrade(*decimal.New(1, 0), *decimal.New(3, 0)) // No Diff
	candle.AddTrade(*decimal.New(1, 0), *decimal.New(3, 0)) // Close

	f, _ := candle.OpenPrice.Float64()
	assert.EqualValues(t, 2, f)
	f, _ = candle.MaxPrice.Float64()
	assert.EqualValues(t, 5, f)
	f, _ = candle.MinPrice.Float64()
	assert.EqualValues(t, 1, f)
	f, _ = candle.ClosePrice.Float64()
	assert.EqualValues(t, 3, f)
	f, _ = candle.Volume.Float64()
	assert.EqualValues(t, 5, f)
	assert.EqualValues(t, 5, candle.TradeCount)
}

func TestCandle_String(t *testing.T) {
	now := time.Now()
	candle := NewCandle(TimePeriod{
		Start: now,
		End:   now.Add(time.Minute),
	})

	candle.ClosePrice = *decimal.New(1, 0)
	candle.OpenPrice = *decimal.New(2, 0)
	candle.MaxPrice = *decimal.New(3, 0)
	candle.MinPrice = decimal.Big{}
	candle.Volume = *decimal.New(10, 0)

	expected := strings.TrimSpace(fmt.Sprintf(`
Time:	%s
Open:	2
Close:	1
High:	3
Low:	0
Volume:	10
`, candle.Period))

	assert.EqualValues(t, expected, candle.String())
}
