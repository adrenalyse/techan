package techan

import (
	"github.com/ericlagergren/decimal"
)

type keltnerChannelIndicator struct {
	ema    Indicator
	atr    Indicator
	mul    decimal.Big
	window int
}

func NewKeltnerChannelUpperIndicator(series *TimeSeries, window int) Indicator {
	return keltnerChannelIndicator{
		atr:    NewAverageTrueRangeIndicator(series, window),
		ema:    NewEMAIndicator(NewClosePriceIndicator(series), window),
		mul:    *decimal.New(1, 0),
		window: window,
	}
}

func NewKeltnerChannelLowerIndicator(series *TimeSeries, window int) Indicator {
	return keltnerChannelIndicator{
		atr:    NewAverageTrueRangeIndicator(series, window),
		ema:    NewEMAIndicator(NewClosePriceIndicator(series), window),
		mul:    *decimal.New(-1, 0),
		window: window,
	}
}

func (kci keltnerChannelIndicator) Calculate(index int) *decimal.Big {
	if index <= kci.window-1 {
		return &decimal.Big{}
	}

	coefficient := decimal.New(2, 0)
	coefficient.Mul(coefficient, &kci.mul)

	tmp := kci.ema.Calculate(index)

	return new(decimal.Big).Add(tmp, coefficient.Mul(kci.atr.Calculate(index), coefficient))
}
