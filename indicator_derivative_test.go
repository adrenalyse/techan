package techan

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDerivativeIndicator(t *testing.T) {
	series := mockTimeSeries("1", "1", "2", "3", "5", "8", "13")
	indicator := DerivativeIndicator{
		Indicator: NewClosePriceIndicator(series),
	}

	t.Run("returns zero at index zero", func(t *testing.T) {
		c := indicator.Calculate(0)
		assert.EqualValues(t, "0", c.String())
	})

	t.Run("returns the derivative", func(t *testing.T) {
		c := indicator.Calculate(1)
		assert.EqualValues(t, "0", c.String())

		for i := 2; i < len(series.Candles); i++ {
			expected := series.Candles[i-2].ClosePrice

			c := indicator.Calculate(i)
			assert.EqualValues(t, expected.String(), c.String())
		}
	})
}
