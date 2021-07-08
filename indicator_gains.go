package techan

import (
	"github.com/ericlagergren/decimal"
)

type gainLossIndicator struct {
	Indicator
	coefficient decimal.Big
}

// NewGainIndicator returns a derivative indicator that returns the gains in the underlying indicator in the last bar,
// if any. If the delta is negative, zero is returned
func NewGainIndicator(indicator Indicator) Indicator {
	return gainLossIndicator{
		Indicator:   indicator,
		coefficient: *decimal.New(1, 0),
	}
}

// NewLossIndicator returns a derivative indicator that returns the losses in the underlying indicator in the last bar,
// if any. If the delta is positive, zero is returned
func NewLossIndicator(indicator Indicator) Indicator {
	return gainLossIndicator{
		Indicator:   indicator,
		coefficient: *decimal.New(-1, 0),
	}
}

func (gli gainLossIndicator) Calculate(index int) *decimal.Big {
	if index == 0 {
		return &decimal.Big{}
	}

	tmp := gli.Indicator.Calculate(index)
	tmp1 := tmp.Sub(tmp, gli.Indicator.Calculate(index-1))

	delta := tmp1.Mul(tmp1, &gli.coefficient)
	if delta.Cmp(&decimal.Big{}) == 1 {
		return delta
	}

	return &decimal.Big{}
}

type cumulativeIndicator struct {
	Indicator
	window int
	mult   decimal.Big
}

// NewCumulativeGainsIndicator returns a derivative indicator which returns all gains made in a base indicator for a given
// window.
func NewCumulativeGainsIndicator(indicator Indicator, window int) Indicator {
	return cumulativeIndicator{
		Indicator: indicator,
		window:    window,
		mult:      *decimal.New(1, 0),
	}
}

// NewCumulativeLossesIndicator returns a derivative indicator which returns all losses in a base indicator for a given
// window.
func NewCumulativeLossesIndicator(indicator Indicator, window int) Indicator {
	return cumulativeIndicator{
		Indicator: indicator,
		window:    window,
		mult:      *decimal.New(-1, 0),
	}
}

func (ci cumulativeIndicator) Calculate(index int) *decimal.Big {
	total := &decimal.Big{}
	tmp := &decimal.Big{}
	for i := Max(1, index-(ci.window-1)); i <= index; i++ {
		diff := tmp.Sub(ci.Indicator.Calculate(i), ci.Indicator.Calculate(i-1))
		if diff.Mul(diff, &ci.mult).Cmp(&decimal.Big{}) == 1 {
			total.Add(total, diff.Abs(diff))
		}
	}

	return total
}

type percentChangeIndicator struct {
	Indicator
}

// NewPercentChangeIndicator returns a derivative indicator which returns the percent change (positive or negative)
// made in a base indicator up until the given indicator
func NewPercentChangeIndicator(indicator Indicator) Indicator {
	return percentChangeIndicator{indicator}
}

func (pgi percentChangeIndicator) Calculate(index int) *decimal.Big {
	if index == 0 {
		return &decimal.Big{}
	}

	cp := pgi.Indicator.Calculate(index)
	cplast := pgi.Indicator.Calculate(index - 1)
	tmp := cp.Quo(cp, cplast)
	return tmp.Sub(tmp, decimal.New(1, 0))
}
