package techan

import (
	"github.com/ericlagergren/decimal"
)

type relativeStrengthIndexIndicator struct {
	rsIndicator Indicator
	oneHundred  decimal.Big
}

// NewRelativeStrengthIndexIndicator returns a derivative Indicator which returns the relative strength index of the base indicator
// in a given time frame. A more in-depth explanation of relative strength index can be found here:
// https://www.investopedia.com/terms/r/rsi.asp
func NewRelativeStrengthIndexIndicator(indicator Indicator, timeframe int) Indicator {
	return relativeStrengthIndexIndicator{
		rsIndicator: NewRelativeStrengthIndicator(indicator, timeframe),
		oneHundred:  *decimal.New(100, 0),
	}
}

func (rsi relativeStrengthIndexIndicator) Calculate(index int) decimal.Big {
	relativeStrength := rsi.rsIndicator.Calculate(index)

	tmp := decimal.New(1, 0)

	return *tmp.Sub(&rsi.oneHundred, tmp.Quo(&rsi.oneHundred, tmp.Add(tmp, &relativeStrength)))
}

type relativeStrengthIndicator struct {
	avgGain Indicator
	avgLoss Indicator
	window  int
}

// NewRelativeStrengthIndicator returns a derivative Indicator which returns the relative strength of the base indicator
// in a given time frame. Relative strength is the average again of up periods during the time frame divided by the
// average loss of down period during the same time frame
func NewRelativeStrengthIndicator(indicator Indicator, timeframe int) Indicator {
	return relativeStrengthIndicator{
		avgGain: NewMMAIndicator(NewGainIndicator(indicator), timeframe),
		avgLoss: NewMMAIndicator(NewLossIndicator(indicator), timeframe),
		window:  timeframe,
	}
}

func (rs relativeStrengthIndicator) Calculate(index int) decimal.Big {
	if index < rs.window-1 {
		return decimal.Big{}
	}

	avgGain := rs.avgGain.Calculate(index)
	avgLoss := rs.avgLoss.Calculate(index)

	if avgLoss.Cmp(&decimal.Big{}) == 0 {
		return *new(decimal.Big).SetInf(false)
	}

	return *avgGain.Quo(&avgGain, &avgLoss)
}
