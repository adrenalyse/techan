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

func (di differenceIndicator) Calculate(index int) decimal.Big {
	tmp1 := di.minuend.Calculate(index)
	tmp2 := di.subtrahend.Calculate(index)
	return *new(decimal.Big).Sub(&tmp1, &tmp2)
}
