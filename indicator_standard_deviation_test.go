package techan

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStandardDeviationIndicator(t *testing.T) {
	t.Run("when index is less than 1, returns 0", func(t *testing.T) {
		series := mockTimeSeries("0", "10")
		stdDev := NewStandardDeviationIndicator(NewClosePriceIndicator(series))

		assert.EqualValues(t, "0", stdDev.Calculate(0).String())
	})

	t.Run("returns the standard deviation when index > 2", func(t *testing.T) {
		series := mockTimeSeriesFl(
			10,
			2,
			38,
			23,
			38,
			23,
			21)

		stdDev := NewStandardDeviationIndicator(NewClosePriceIndicator(series))

		assert.EqualValues(t, "4", stdDev.Calculate(1).String())
		assert.EqualValues(t, "15.43444920372030", stdDev.Calculate(2).String())
		assert.EqualValues(t, "13.64505404899519", stdDev.Calculate(3).String())
		assert.EqualValues(t, "14.53822547630900", stdDev.Calculate(4).String())
		assert.EqualValues(t, "13.27487183449325", stdDev.Calculate(5).String())
		assert.EqualValues(t, "12.29899614287479", stdDev.Calculate(6).String())
	})
}
