package techan

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTrendIndicator(t *testing.T) {
	t.Run("returns the correct slope of the trend", func(t *testing.T) {
		tests := []struct {
			series         []float64
			expectedResult string
		}{
			{
				series:         []float64{0, 1, 2, 3},
				expectedResult: "1",
			},
			{
				series:         []float64{0, 2, 4, 6},
				expectedResult: "2",
			},
			{
				series:         []float64{5, 4, 3, 2},
				expectedResult: "-1",
			},
		}

		for _, test := range tests {
			series := mockTimeSeriesFl(test.series...)
			indicator := NewTrendlineIndicator(NewClosePriceIndicator(series), 4)

			c := indicator.Calculate(3)
			assert.EqualValues(t, test.expectedResult, c.String())
		}
	})

	t.Run("respects the window", func(t *testing.T) {
		series := mockTimeSeriesFl(-100, 1000, 0, 1, 2, 3)
		indicator := NewTrendlineIndicator(NewClosePriceIndicator(series), 4)
		c := indicator.Calculate(5)
		assert.EqualValues(t, "1", c.String())
	})

	t.Run("does not allow an index out of bounds on the low end", func(t *testing.T) {
		series := mockTimeSeriesFl(0, 1)
		indicator := NewTrendlineIndicator(NewClosePriceIndicator(series), 4)
		c := indicator.Calculate(1)
		assert.EqualValues(t, "1", c.String())
	})
}
