package techan

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStandardDeviationIndicator(t *testing.T) {
	t.Run("when index is less than 1, returns 0", func(t *testing.T) {
		series := mockTimeSeries("0", "10")
		stdDev := NewStandardDeviationIndicator(NewClosePriceIndicator(series))

		c := stdDev.Calculate(0)
		assert.EqualValues(t, "0", c.String())
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

		c := stdDev.Calculate(1)
		assert.EqualValues(t, "4", c.String())
		c = stdDev.Calculate(2)
		assert.EqualValues(t, "15.43444920372030", c.String())
		c = stdDev.Calculate(3)
		assert.EqualValues(t, "13.64505404899519", c.String())
		c = stdDev.Calculate(4)
		assert.EqualValues(t, "14.53822547630900", c.String())
		c = stdDev.Calculate(5)
		assert.EqualValues(t, "13.27487183449325", c.String())
		c = stdDev.Calculate(6)
		assert.EqualValues(t, "12.29899614287479", c.String())
	})
}
