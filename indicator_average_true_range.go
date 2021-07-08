package techan

import (
	"github.com/sdcoffey/big"
	"log"
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

func (atr averageTrueRangeIndicator) Calculate(index int) big.Decimal {
	if index < atr.window {
		return big.ZERO
	}

	sum := big.ZERO
	indicator := NewTrueRangeIndicator(atr.series)

	for i := index; i > index-atr.window; i-- {
		//log.Println(i, sum, indicator.Calculate(i))
		sum = sum.Add(indicator.Calculate(i))
	}

	r := sum.Div(big.NewFromInt(atr.window))
	log.Println(sum, atr.window, r)

	return r
}
