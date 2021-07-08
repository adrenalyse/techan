package techan

import (
	"github.com/ericlagergren/decimal"
)

type trendLineIndicator struct {
	indicator Indicator
	window    int
}

// NewTrendlineIndicator returns an indicator whose output is the slope of the trend
// line given by the values in the window.
func NewTrendlineIndicator(indicator Indicator, window int) Indicator {
	return trendLineIndicator{
		indicator: indicator,
		window:    window,
	}
}

func (tli trendLineIndicator) Calculate(index int) *decimal.Big {
	window := Min(index+1, tli.window)

	values := make([]*decimal.Big, window)

	for i := 0; i < window; i++ {
		values[i] = tli.indicator.Calculate(index - (window - 1) + i)
	}

	tmp1 := decimal.New(1, 0)
	tmp2 := &decimal.Big{}
	n := tmp1.Mul(tmp1, tmp2.SetFloat64(float64(window))) // tmp1 occupé

	tmp3 := sumX(values)
	tmp4 := sumY(values)
	ab := tmp2.Sub(tmp2.Mul(sumXy(values), n), tmp4.Mul(tmp3, tmp4)) // tmp1, tmp2, tmp3 occupés
	cd := tmp1.Sub(tmp1.Mul(sumX2(values), n), tmp3.Mul(tmp3, tmp3)) // tmp1, tmp2 occupés

	return ab.Quo(ab, cd)
}

func sumX(decimals []*decimal.Big) (s *decimal.Big) {
	s = &decimal.Big{}
	tmp := &decimal.Big{}
	for i := range decimals {
		s.Add(s, tmp.SetFloat64(float64(i)))
	}

	return s
}

func sumY(decimals []*decimal.Big) (b *decimal.Big) {
	b = &decimal.Big{}
	for _, d := range decimals {
		b.Add(b, d)
	}

	return b
}

func sumXy(decimals []*decimal.Big) (b *decimal.Big) {
	b = &decimal.Big{}
	tmp := &decimal.Big{}
	for i, d := range decimals {
		b.Add(b, tmp.Mul(d, tmp.SetFloat64(float64(i))))
	}

	return b
}

func sumX2(decimals []*decimal.Big) *decimal.Big {
	b := &decimal.Big{}
	tmp := &decimal.Big{}
	for i := range decimals {
		tmp.SetFloat64(float64(i))
		b.Add(b, tmp.Mul(tmp, tmp))
	}

	return b
}
