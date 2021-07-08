package techan

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVarianceIndicator(t *testing.T) {
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

		varInd := NewVarianceIndicator(NewClosePriceIndicator(series))

		c := varInd.Calculate(1)
		assert.EqualValues(t, "16", c.String())
		c = varInd.Calculate(2)
		assert.EqualValues(t, "238.2222222222222", c.String())
		c = varInd.Calculate(3)
		assert.EqualValues(t, "186.1875", c.String())
		c = varInd.Calculate(4)
		assert.EqualValues(t, "211.36", c.String())
		c = varInd.Calculate(5)
		assert.EqualValues(t, "176.2222222222222", c.String())
		c = varInd.Calculate(6)
		assert.EqualValues(t, "151.265306122449", c.String())
	})
}
