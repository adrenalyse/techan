package techan

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVarianceIndicator(t *testing.T) {
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

		varInd := NewVarianceIndicator(NewClosePriceIndicator(series))

		assert.EqualValues(t, "16", varInd.Calculate(1).String())
		assert.EqualValues(t, "238.2222222222222", varInd.Calculate(2).String())
		assert.EqualValues(t, "186.1875", varInd.Calculate(3).String())
		assert.EqualValues(t, "211.36", varInd.Calculate(4).String())
		assert.EqualValues(t, "176.2222222222222", varInd.Calculate(5).String())
		assert.EqualValues(t, "151.265306122449", varInd.Calculate(6).String())
	})
}
