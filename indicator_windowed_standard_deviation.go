package techan

import (
	"github.com/ericlagergren/decimal"
	"github.com/ericlagergren/decimal/math"
)

type windowedStandardDeviationIndicator struct {
	Indicator
	movingAverage Indicator
	window        int
}

// NewWindowedStandardDeviationIndicator returns a indicator which calculates the standard deviation of the underlying
// indicator over a window
func NewWindowedStandardDeviationIndicator(ind Indicator, window int) Indicator {
	return windowedStandardDeviationIndicator{
		Indicator:     ind,
		movingAverage: NewSimpleMovingAverage(ind, window),
		window:        window,
	}
}

func (sdi windowedStandardDeviationIndicator) Calculate(index int) decimal.Big {
	avg := sdi.movingAverage.Calculate(index)
	variance := &decimal.Big{}
	var tmp decimal.Big

	for i := Max(0, index-sdi.window+1); i <= index; i++ {
		tmp = sdi.Indicator.Calculate(i)
		tmp.Sub(&tmp, &avg)
		pow := tmp.Mul(&tmp, &tmp)
		variance.Add(variance, pow)
	}
	realwindow := Min(sdi.window, index+1)

	variance.Quo(variance, avg.SetFloat64(float64(realwindow)))
	return *math.Sqrt(variance, variance)
}
