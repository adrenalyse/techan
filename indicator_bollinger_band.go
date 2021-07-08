package techan

import (
	"github.com/ericlagergren/decimal"
)

type bbandIndicator struct {
	ma     Indicator
	stdev  Indicator
	muladd decimal.Big
}

// NewBollingerUpperBandIndicator a a derivative indicator which returns the upper bound of a bollinger band
// on the underlying indicator
func NewBollingerUpperBandIndicator(indicator Indicator, window int, sigma float64) Indicator {
	return bbandIndicator{
		ma:     NewSimpleMovingAverage(indicator, window),
		stdev:  NewWindowedStandardDeviationIndicator(indicator, window),
		muladd: *new(decimal.Big).SetFloat64(sigma),
	}
}

// NewBollingerLowerBandIndicator returns a a derivative indicator which returns the lower bound of a bollinger band
// on the underlying indicator
func NewBollingerLowerBandIndicator(indicator Indicator, window int, sigma float64) Indicator {
	return bbandIndicator{
		ma:     NewSimpleMovingAverage(indicator, window),
		stdev:  NewWindowedStandardDeviationIndicator(indicator, window),
		muladd: *new(decimal.Big).SetFloat64(-sigma),
	}
}

func (bbi bbandIndicator) Calculate(index int) *decimal.Big {
	tmp := bbi.ma.Calculate(index)
	// bbi.ma.Calculate(index) + bbi.stdev.Calculate(index) * bbi.muladd
	return tmp.FMA(bbi.stdev.Calculate(index), &bbi.muladd, tmp)
}
