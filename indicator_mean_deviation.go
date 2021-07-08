package techan

import (
	"github.com/ericlagergren/decimal"
)

type meanDeviationIndicator struct {
	Indicator
	movingAverage Indicator
	window        int
}

// NewMeanDeviationIndicator returns a derivative Indicator which returns the mean deviation of a base indicator
// in a given window. Mean deviation is an average of all values on the base indicator from the mean of that indicator.
func NewMeanDeviationIndicator(indicator Indicator, window int) Indicator {
	return meanDeviationIndicator{
		Indicator:     indicator,
		movingAverage: NewSimpleMovingAverage(indicator, window),
		window:        window,
	}
}

func (mdi meanDeviationIndicator) Calculate(index int) decimal.Big {
	if index < mdi.window-1 {
		return decimal.Big{}
	}

	average := mdi.movingAverage.Calculate(index)
	start := Max(0, index-mdi.window+1)
	absoluteDeviations := &decimal.Big{}

	for i := start; i <= index; i++ {
		tmp := mdi.Indicator.Calculate(i)
		absoluteDeviations.Add(absoluteDeviations, tmp.Abs(tmp.Sub(&average, &tmp)))
	}

	return *absoluteDeviations.Quo(absoluteDeviations, average.SetFloat64(float64(Min(mdi.window, index-start+1))))
}
