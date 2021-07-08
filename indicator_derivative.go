package techan

import (
	"github.com/ericlagergren/decimal"
)

// DerivativeIndicator returns an indicator that calculates the derivative of the underlying Indicator.
// The derivative is defined as the difference between the value at the previous index and the value at the current index.
// Eg series [1, 1, 2, 3, 5, 8] -> [0, 0, 1, 1, 2, 3]
type DerivativeIndicator struct {
	Indicator Indicator
}

// Calculate returns the derivative of the underlying indicator. At index 0, it will always return 0.
func (di DerivativeIndicator) Calculate(index int) decimal.Big {
	if index == 0 {
		return decimal.Big{}
	}

	tmp1 := di.Indicator.Calculate(index)
	tmp2 := di.Indicator.Calculate(index - 1)
	return *new(decimal.Big).Sub(&tmp1, &tmp2)
}
