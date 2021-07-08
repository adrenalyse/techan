package techan

import "testing"

func TestAverageTrueRangeIndicator(t *testing.T) {
	atrIndicator := NewAverageTrueRangeIndicator(mockedTimeSeries, 3)

	expectedValues := []float64{
		0,
		0,
		0,
		2,
		2,
		2,
		2,
		2,
		2,
		2,
		2.346666666666667,
		2.346666666666667,
	}

	indicatorEquals(t, expectedValues, atrIndicator)
}
