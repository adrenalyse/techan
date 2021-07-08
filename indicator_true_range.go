package techan

import (
	"github.com/ericlagergren/decimal"
)

type trueRangeIndicator struct {
	series *TimeSeries
}

// NewTrueRangeIndicator returns a base indicator
// which calculates the true range at the current point in time for a series
// https://www.investopedia.com/terms/a/atr.asp
func NewTrueRangeIndicator(series *TimeSeries) Indicator {
	return trueRangeIndicator{
		series: series,
	}
}

func (tri trueRangeIndicator) Calculate(index int) *decimal.Big {
	if index-1 < 0 {
		return &decimal.Big{}
	}

	candle := tri.series.Candles[index]
	previousClose := tri.series.Candles[index-1].ClosePrice

	trueHigh := previousClose
	if candle.MaxPrice.Cmp(previousClose) == 1 {
		trueHigh = candle.MaxPrice
	}

	trueLow := previousClose
	if candle.MinPrice.Cmp(previousClose) == -1 {
		trueLow = candle.MinPrice
	}

	return new(decimal.Big).Sub(trueHigh, trueLow)
}
