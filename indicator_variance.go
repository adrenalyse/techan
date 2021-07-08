package techan

import (
	"github.com/ericlagergren/decimal"
)

// NewVarianceIndicator provides a way to find the variance in a base indicator, where variances is the sum of squared
// deviations from the mean at any given index in the time series.
func NewVarianceIndicator(ind Indicator) Indicator {
	return varianceIndicator{
		Indicator: ind,
	}
}

type varianceIndicator struct {
	Indicator Indicator
}

// Calculate returns the Variance for this indicator at the given index
func (vi varianceIndicator) Calculate(index int) decimal.Big {
	if index < 1 {
		return decimal.Big{}
	}

	avgIndicator := NewSimpleMovingAverage(vi.Indicator, index+1)
	avg := avgIndicator.Calculate(index)
	variance := &decimal.Big{}
	var tmp decimal.Big

	for i := 0; i <= index; i++ {
		tmp = vi.Indicator.Calculate(i)
		tmp.Sub(&tmp, &avg)
		pow := tmp.Mul(&tmp, &tmp)
		variance.Add(variance, pow)
	}

	return *variance.Quo(variance, avg.SetFloat64(float64(index+1)))
}
