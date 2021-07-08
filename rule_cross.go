package techan

// NewCrossUpIndicatorRule returns a new rule that is satisfied when the lower indicator has crossed above the upper
// indicator.
func NewCrossUpIndicatorRule(upper, lower Indicator) Rule {
	return crossRule{
		upper: upper,
		lower: lower,
		cmp:   1,
	}
}

// NewCrossDownIndicatorRule returns a new rule that is satisfied when the upper indicator has crossed below the lower
// indicator.
func NewCrossDownIndicatorRule(upper, lower Indicator) Rule {
	return crossRule{
		upper: lower,
		lower: upper,
		cmp:   -1,
	}
}

type crossRule struct {
	upper Indicator
	lower Indicator
	cmp   int
}

func (cr crossRule) IsSatisfied(index int, record *TradingRecord) bool {
	i := index

	if i == 0 {
		return false
	}

	tmp := cr.upper.Calculate(i)
	tmpp := cr.lower.Calculate(i)
	if cmp := tmpp.Cmp(&tmp); cmp == 0 || cmp == cr.cmp {
		for ; i >= 0; i-- {
			tmp1 := cr.upper.Calculate(i)
			tmp2 := cr.lower.Calculate(i)
			if cmp = tmp2.Cmp(&tmp1); cmp == 0 || cmp == -cr.cmp {
				return true
			}
		}
	}

	return false
}
