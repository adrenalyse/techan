package techan

import (
	"github.com/ericlagergren/decimal"
)

type stopLossRule struct {
	Indicator
	tolerance decimal.Big
}

// NewStopLossRule returns a new rule that is satisfied when the given loss tolerance (a percentage) is met or exceeded.
// Loss tolerance should be a value between -1 and 1.
func NewStopLossRule(series *TimeSeries, lossTolerance float64) Rule {
	return stopLossRule{
		Indicator: NewClosePriceIndicator(series),
		tolerance: *new(decimal.Big).SetFloat64(lossTolerance),
	}
}

func (slr stopLossRule) IsSatisfied(index int, record *TradingRecord) bool {
	if !record.CurrentPosition().IsOpen() {
		return false
	}

	openPrice := record.CurrentPosition().CostBasis()
	tmp := slr.Indicator.Calculate(index)
	loss := tmp.Sub(tmp.Quo(&tmp, openPrice), decimal.New(1, 0))
	comp := loss.Cmp(&slr.tolerance)
	return comp == -1 || comp == 0
}
