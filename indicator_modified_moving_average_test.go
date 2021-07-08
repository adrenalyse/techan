package techan

import (
	"testing"
)

func TestModifiedMovingAverage(t *testing.T) {
	indicator := NewMMAIndicator(NewClosePriceIndicator(mockedTimeSeries), 3)

	expected := []float64{
		0,
		0,
		64.09,
		63.97,
		63.83,
		63.61666666666667,
		63.71444444444445,
		63.75962962962963,
		63.48975308641975,
		63.4498353909465,
		62.74322359396433,
		62.33214906264289,
	}

	indicatorEquals(t, expected, indicator)
}
