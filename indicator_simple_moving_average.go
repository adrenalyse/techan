package techan

import (
	"github.com/ericlagergren/decimal"
)

type smaIndicator struct {
	indicator Indicator
	window    int
}

// NewSimpleMovingAverage returns a derivative Indicator which returns the average of the current value and preceding
// values in the given windowSize.
func NewSimpleMovingAverage(indicator Indicator, window int) Indicator {
	return smaIndicator{indicator, window}
}

func (sma smaIndicator) Calculate(index int) *decimal.Big {
	if index < sma.window-1 {
		return &decimal.Big{}
	}

	sum := &decimal.Big{}
	tmp := &decimal.Big{}
	for i := index; i > index-sma.window; i-- {
		tmp = sma.indicator.Calculate(i)
		sum.Add(sum, tmp)
	}

	return sum.Quo(sum, tmp.SetUint64(uint64(sma.window)))
}
