package techan

import (
	"github.com/ericlagergren/decimal"
)

type averageTrueRangeIndicator struct {
	series *TimeSeries
	window int
}

// NewAverageTrueRangeIndicator returns a base indicator that calculates the average true range of the
// underlying over a window
// https://www.investopedia.com/terms/a/atr.asp
func NewAverageTrueRangeIndicator(series *TimeSeries, window int) Indicator {
	return averageTrueRangeIndicator{
		series: series,
		window: window,
	}
}

func (atr averageTrueRangeIndicator) Calculate(index int) *decimal.Big {
	if index < atr.window {
		return &decimal.Big{}
	}

	sum := &decimal.Big{}
	indicator := NewTrueRangeIndicator(atr.series)

	for i := index; i > index-atr.window; i-- {
		sum.Add(sum, indicator.Calculate(i))
	}

	return sum.Quo(sum, decimal.New(int64(atr.window), 0))
}
