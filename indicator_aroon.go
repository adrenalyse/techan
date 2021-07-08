package techan

import (
	"github.com/ericlagergren/decimal"
	"math"
)

type aroonIndicator struct {
	indicator Indicator
	window    int
	direction decimal.Big
	lowIndex  int
}

func (ai *aroonIndicator) Calculate(index int) decimal.Big {
	if index < ai.window-1 {
		return decimal.Big{}
	}

	oneHundred := decimal.New(100, 0)
	ai.lowIndex = ai.findLowIndex(index)
	// pSince
	tmp1 := new(decimal.Big).SetFloat64(float64(index - ai.lowIndex))
	// windowAsDecimal
	tmp2 := new(decimal.Big).SetFloat64(float64(ai.window))

	return *tmp2.Mul(tmp1.Quo(tmp1.Sub(tmp2, tmp1), tmp2), oneHundred)
}

func (ai aroonIndicator) findLowIndex(index int) int {
	if ai.lowIndex < 1 || ai.lowIndex < index-ai.window {
		lv := new(decimal.Big).SetFloat64(math.MaxFloat64)
		lowIndex := -1
		for i := (index + 1) - ai.window; i <= index; i++ {
			tmp := ai.indicator.Calculate(i)
			value := tmp.Mul(&tmp, &ai.direction)
			if value.Cmp(lv) == -1 {
				lv = value
				lowIndex = i
			}
		}

		return lowIndex
	}

	tmp1 := ai.indicator.Calculate(index)
	v1 := tmp1.Mul(&tmp1, &ai.direction)
	tmp2 := ai.indicator.Calculate(ai.lowIndex)
	v2 := tmp2.Mul(&tmp2, &ai.direction)

	if v1.Cmp(v2) == -1 {
		return index
	}

	return ai.lowIndex
}

// NewAroonUpIndicator returns a derivative indicator that will return a value based on
// the number of ticks since the highest price in the window
// https://www.investopedia.com/terms/a/aroon.asp
//
// Note: this indicator should be constructed with a either a HighPriceIndicator or a derivative thereof
func NewAroonUpIndicator(indicator Indicator, window int) Indicator {
	return &aroonIndicator{
		indicator: indicator,
		window:    window,
		direction: *decimal.New(-1, 0),
		lowIndex:  -1,
	}
}

// NewAroonDownIndicator returns a derivative indicator that will return a value based on
// the number of ticks since the lowest price in the window
// https://www.investopedia.com/terms/a/aroon.asp
//
// Note: this indicator should be constructed with a either a LowPriceIndicator or a derivative thereof
func NewAroonDownIndicator(indicator Indicator, window int) Indicator {
	return &aroonIndicator{
		indicator: indicator,
		window:    window,
		direction: *decimal.New(1, 0),
		lowIndex:  -1,
	}
}
