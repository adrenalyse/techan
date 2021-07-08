package techan

import (
	"fmt"
	"github.com/ericlagergren/decimal"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewVolumeIndicator(t *testing.T) {
	assert.NotNil(t, NewVolumeIndicator(NewTimeSeries()))
}

func TestVolumeIndicator_Calculate(t *testing.T) {
	series := NewTimeSeries()

	candle := NewCandle(TimePeriod{
		Start: time.Now(),
		End:   time.Now().Add(time.Minute),
	})
	candle.Volume = *decimal.New(12080, 4)

	series.AddCandle(candle)

	indicator := NewVolumeIndicator(series)
	c := indicator.Calculate(0)
	assert.EqualValues(t, "1.2080", c.String())
}

func TestTypicalPriceIndicator_Calculate(t *testing.T) {
	series := NewTimeSeries()

	candle := NewCandle(TimePeriod{
		Start: time.Now(),
		End:   time.Now().Add(time.Minute),
	})
	candle.MinPrice = *decimal.New(12080, 4)
	candle.MaxPrice = *decimal.New(122, 2)
	candle.ClosePrice = *decimal.New(1215, 3)

	series.AddCandle(candle)

	typicalPrice := NewTypicalPriceIndicator(series).Calculate(0)
	typicalPrice = NewTypicalPriceIndicator(series).Calculate(0)

	f, _ := typicalPrice.Float64()
	assert.EqualValues(t, "1.214333", fmt.Sprintf("%f", f))
}
