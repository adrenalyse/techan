package techan

import (
	"github.com/ericlagergren/decimal"
)

// NewMaximumValueIndicator returns a derivative Indicator which returns the maximum value
// present in a given window. Use a window value of -1 to include all values in the
// underlying indicator.
func NewMaximumValueIndicator(ind Indicator, window int) Indicator {
	return maximumValueIndicator{
		indicator: ind,
		window:    window,
	}
}

type maximumValueIndicator struct {
	indicator Indicator
	window    int
}

func (mvi maximumValueIndicator) Calculate(index int) decimal.Big {
	maxValue := new(decimal.Big).SetInf(true)

	start := 0
	if mvi.window > 0 {
		start = Max(index-mvi.window+1, 0)
	}

	for i := start; i <= index; i++ {
		value := mvi.indicator.Calculate(i)
		if value.Cmp(maxValue) == 1 {
			maxValue = &value
		}
	}

	return *maxValue
}
