package techan

import (
	"github.com/ericlagergren/decimal"
)

type differenceIndicator struct {
	minuend    Indicator
	subtrahend Indicator
}

// NewDifferenceIndicator returns an indicator which returns the difference between one indicator (minuend) and a second
// indicator (subtrahend).
func NewDifferenceIndicator(minuend, subtrahend Indicator) Indicator {
	return differenceIndicator{
		minuend:    minuend,
		subtrahend: subtrahend,
	}
}

func (di differenceIndicator) Calculate(index int) *decimal.Big {
	return new(decimal.Big).Sub(di.minuend.Calculate(index), di.subtrahend.Calculate(index))
}
