package techan

import (
	"testing"
)

func TestKeltnerChannel(t *testing.T) {
	t.Run("Upper", func(t *testing.T) {
		upper := NewKeltnerChannelUpperIndicator(mockedTimeSeries, 3)

		expectedValues := []float64{
			0,
			0,
			0,
			67.91,
			67.73,
			67.46,
			67.685,
			67.7675,
			67.35875,
			67.364375,
			67.04052083333333,
			66.62192708333333,
		}

		indicatorEquals(t, expectedValues, upper)
	})

	t.Run("Lower", func(t *testing.T) {
		lower := NewKeltnerChannelLowerIndicator(mockedTimeSeries, 3)

		expectedValues := []float64{
			0,
			0,
			0,
			59.91,
			59.73,
			59.46,
			59.685,
			59.7675,
			59.35875,
			59.364375,
			57.65385416666667,
			57.23526041666667,
		}

		indicatorEquals(t, expectedValues, lower)
	})
}
