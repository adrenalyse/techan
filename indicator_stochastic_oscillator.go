package techan

import (
	"github.com/ericlagergren/decimal"
)

type kIndicator struct {
	closePrice Indicator
	minValue   Indicator
	maxValue   Indicator
	window     int
}

// NewFastStochasticIndicator returns a derivative Indicator which returns the fast stochastic indicator (%K) for the
// given window.
// https://www.investopedia.com/terms/s/stochasticoscillator.asp
func NewFastStochasticIndicator(series *TimeSeries, timeframe int) Indicator {
	return kIndicator{
		closePrice: NewClosePriceIndicator(series),
		minValue:   NewMinimumValueIndicator(NewLowPriceIndicator(series), timeframe),
		maxValue:   NewMaximumValueIndicator(NewHighPriceIndicator(series), timeframe),
		window:     timeframe,
	}
}

func (k kIndicator) Calculate(index int) decimal.Big {
	closeVal := k.closePrice.Calculate(index)
	minVal := k.minValue.Calculate(index)
	maxVal := k.maxValue.Calculate(index)

	if minVal.Cmp(&maxVal) == 0 {
		return *new(decimal.Big).SetInf(false)
	}
	r := decimal.New(100, 0)
	return *r.Mul(maxVal.Quo(closeVal.Sub(&closeVal, &minVal), maxVal.Sub(&maxVal, &minVal)), r)
}

type dIndicator struct {
	k      Indicator
	window int
}

// NewSlowStochasticIndicator returns a derivative Indicator which returns the slow stochastic indicator (%D) for the
// given window.
// https://www.investopedia.com/terms/s/stochasticoscillator.asp
func NewSlowStochasticIndicator(k Indicator, window int) Indicator {
	return dIndicator{k, window}
}

func (d dIndicator) Calculate(index int) decimal.Big {
	return NewSimpleMovingAverage(d.k, d.window).Calculate(index)
}
