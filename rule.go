package techan

import (
	"github.com/ericlagergren/decimal"
)

// Rule is an interface describing an algorithm by which a set of criteria may be satisfied
type Rule interface {
	IsSatisfied(index int, record *TradingRecord) bool
}

// And returns a new rule whereby BOTH of the passed-in rules must be satisfied for the rule to be satisfied
func And(r1, r2 Rule) Rule {
	return andRule{r1, r2}
}

// Or returns a new rule whereby ONE OF the passed-in rules must be satisfied for the rule to be satisfied
func Or(r1, r2 Rule) Rule {
	return orRule{r1, r2}
}

type andRule struct {
	r1 Rule
	r2 Rule
}

func (ar andRule) IsSatisfied(index int, record *TradingRecord) bool {
	return ar.r1.IsSatisfied(index, record) && ar.r2.IsSatisfied(index, record)
}

type orRule struct {
	r1 Rule
	r2 Rule
}

func (or orRule) IsSatisfied(index int, record *TradingRecord) bool {
	return or.r1.IsSatisfied(index, record) || or.r2.IsSatisfied(index, record)
}

// OverIndicatorRule is a rule where the First Indicator must be greater than the Second Indicator to be Satisfied
type OverIndicatorRule struct {
	First  Indicator
	Second Indicator
}

// IsSatisfied returns true when the First Indicator is greater than the Second Indicator
func (oir OverIndicatorRule) IsSatisfied(index int, record *TradingRecord) bool {
	f := oir.First.Calculate(index)
	s := oir.Second.Calculate(index)
	return f.Cmp(&s) == 1
}

// UnderIndicatorRule is a rule where the First Indicator must be less than the Second Indicator to be Satisfied
type UnderIndicatorRule struct {
	First  Indicator
	Second Indicator
}

// IsSatisfied returns true when the First Indicator is less than the Second Indicator
func (uir UnderIndicatorRule) IsSatisfied(index int, record *TradingRecord) bool {
	f := uir.First.Calculate(index)
	s := uir.Second.Calculate(index)
	return f.Cmp(&s) == -1
}

type percentChangeRule struct {
	indicator Indicator
	percent   decimal.Big
}

func (pgr percentChangeRule) IsSatisfied(index int, record *TradingRecord) bool {
	i := pgr.indicator.Calculate(index)
	return i.Abs(&i).Cmp(pgr.percent.Abs(&pgr.percent)) == 1
}

// NewPercentChangeRule returns a rule whereby the given Indicator must have changed by a given percentage to be satisfied.
// You should specify percent as a float value between -1 and 1
func NewPercentChangeRule(indicator Indicator, percent float64) Rule {
	return percentChangeRule{
		indicator: NewPercentChangeIndicator(indicator),
		percent:   *new(decimal.Big).SetFloat64(percent),
	}
}
