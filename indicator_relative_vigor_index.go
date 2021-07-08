package techan

import (
	"github.com/ericlagergren/decimal"
)

type relativeVigorIndexIndicator struct {
	numerator   Indicator
	denominator Indicator
}

// NewRelativeVigorIndexIndicator returns an Indicator which returns the index of the relative vigor of the prices of
// a sercurity. Relative Vigor Index is simply the difference of the previous four days' close and open prices divided
// by the difference between the previous four days high and low prices. A more in-depth explanation of relative vigor
// index can be found here: https://www.fidelity.com/learning-center/trading-investing/technical-analysis/technical-indicator-guide/relative-vigor-index
func NewRelativeVigorIndexIndicator(series *TimeSeries) Indicator {
	return relativeVigorIndexIndicator{
		numerator:   NewDifferenceIndicator(NewClosePriceIndicator(series), NewOpenPriceIndicator(series)),
		denominator: NewDifferenceIndicator(NewHighPriceIndicator(series), NewLowPriceIndicator(series)),
	}
}

func (rvii relativeVigorIndexIndicator) Calculate(index int) decimal.Big {
	if index < 3 {
		return decimal.Big{}
	}

	two := decimal.New(2, 0)

	a := rvii.numerator.Calculate(index)
	tmp1 := rvii.numerator.Calculate(index - 1)
	b := tmp1.Mul(&tmp1, two)
	tmp2 := rvii.numerator.Calculate(index - 2)
	c := tmp2.Mul(&tmp2, two)
	d := rvii.numerator.Calculate(index - 3)

	num := a.Quo(a.Add(&a, b.Add(b, c.Add(c, &d))), decimal.New(6, 0))

	e := rvii.denominator.Calculate(index)
	tmp1 = rvii.denominator.Calculate(index - 1)
	f := tmp1.Mul(&tmp1, two)
	tmp2 = rvii.denominator.Calculate(index - 2)
	g := tmp2.Mul(&tmp2, two)
	h := rvii.denominator.Calculate(index - 3)

	denom := e.Quo(e.Add(&e, f.Add(f, g.Add(g, &h))), decimal.New(6, 0))

	return *two.Quo(num, denom)
}

type relativeVigorIndexSignalLine struct {
	relativeVigorIndex Indicator
}

// NewRelativeVigorSignalLine returns an Indicator intended to be used in conjunction with Relative vigor index, which
// returns the average value of the last 4 indices of the RVI indicator.
func NewRelativeVigorSignalLine(series *TimeSeries) Indicator {
	return relativeVigorIndexSignalLine{
		relativeVigorIndex: NewRelativeVigorIndexIndicator(series),
	}
}

func (rvsn relativeVigorIndexSignalLine) Calculate(index int) decimal.Big {
	if index < 3 {
		return decimal.Big{}
	}

	rvi := rvsn.relativeVigorIndex.Calculate(index)
	tmp := rvsn.relativeVigorIndex.Calculate(index - 1)
	tmp3 := rvsn.relativeVigorIndex.Calculate(index - 1)
	i := tmp.Mul(&tmp3, decimal.New(2, 0))
	tmp1 := rvsn.relativeVigorIndex.Calculate(index - 2)
	j := tmp1.Mul(&tmp1, decimal.New(2, 0))
	k := rvsn.relativeVigorIndex.Calculate(index - 3)

	r := decimal.New(6, 0)

	return *r.Quo(rvi.Add(&rvi, i.Add(i, j.Add(j, &k))), r)
}
