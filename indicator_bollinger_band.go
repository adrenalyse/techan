package techan

import (
	"github.com/sdcoffey/big"
	"log"
)

type bbandIndicator struct {
	ma     Indicator
	stdev  Indicator
	muladd big.Decimal
}

// NewBollingerUpperBandIndicator a a derivative indicator which returns the upper bound of a bollinger band
// on the underlying indicator
func NewBollingerUpperBandIndicator(indicator Indicator, window int, sigma float64) Indicator {
	return bbandIndicator{
		ma:     NewSimpleMovingAverage(indicator, window),
		stdev:  NewWindowedStandardDeviationIndicator(indicator, window),
		muladd: big.NewDecimal(sigma),
	}
}

// NewBollingerLowerBandIndicator returns a a derivative indicator which returns the lower bound of a bollinger band
// on the underlying indicator
func NewBollingerLowerBandIndicator(indicator Indicator, window int, sigma float64) Indicator {
	return bbandIndicator{
		ma:     NewSimpleMovingAverage(indicator, window),
		stdev:  NewWindowedStandardDeviationIndicator(indicator, window),
		muladd: big.NewDecimal(-sigma),
	}
}

func (bbi bbandIndicator) Calculate(index int) big.Decimal {
	tmp := bbi.ma.Calculate(index)
	tmp1 := bbi.stdev.Calculate(index)
	log.Println(tmp, tmp1)
	return tmp.Add(tmp1.Mul(bbi.muladd))
}
