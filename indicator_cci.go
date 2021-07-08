package techan

import (
	"github.com/ericlagergren/decimal"
)

type commidityChannelIndexIndicator struct {
	series *TimeSeries
	window int
}

// NewCCIIndicator Returns a new Commodity Channel Index Indicator
// http://stockcharts.com/school/doku.php?id=chart_school:technical_indicators:commodity_channel_index_cci
func NewCCIIndicator(ts *TimeSeries, window int) Indicator {
	return commidityChannelIndexIndicator{
		series: ts,
		window: window,
	}
}

func (ccii commidityChannelIndexIndicator) Calculate(index int) decimal.Big {
	typicalPrice := NewTypicalPriceIndicator(ccii.series)
	typicalPriceSma := NewSimpleMovingAverage(typicalPrice, ccii.window)
	meanDeviation := NewMeanDeviationIndicator(NewClosePriceIndicator(ccii.series), ccii.window)

	// (typicalPrice.Calculate(index) - typicalPriceSma.Calculate(index)) / (meanDeviation.Calculate(index) * 0.015)
	tmp := typicalPrice.Calculate(index)
	tmp1 := typicalPriceSma.Calculate(index)
	tmp.Sub(&tmp, &tmp1) // tmp used
	tmp2 := meanDeviation.Calculate(index)
	tmp1.Mul(&tmp2, tmp1.SetFloat64(0.015))

	return *new(decimal.Big).Quo(&tmp, &tmp1)
}
