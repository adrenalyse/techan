package techan

import (
	"github.com/ericlagergren/decimal"
)

type volumeIndicator struct {
	*TimeSeries
}

// NewVolumeIndicator returns an indicator which returns the volume of a candle for a given index
func NewVolumeIndicator(series *TimeSeries) Indicator {
	return volumeIndicator{series}
}

func (vi volumeIndicator) Calculate(index int) *decimal.Big {
	return new(decimal.Big).Copy(vi.Candles[index].Volume)
}

type closePriceIndicator struct {
	*TimeSeries
}

// NewClosePriceIndicator returns an Indicator which returns the close price of a candle for a given index
func NewClosePriceIndicator(series *TimeSeries) Indicator {
	return closePriceIndicator{series}
}

func (cpi closePriceIndicator) Calculate(index int) *decimal.Big {
	return new(decimal.Big).Copy(cpi.Candles[index].ClosePrice)
}

type highPriceIndicator struct {
	*TimeSeries
}

// NewHighPriceIndicator returns an Indicator which returns the high price of a candle for a given index
func NewHighPriceIndicator(series *TimeSeries) Indicator {
	return highPriceIndicator{
		series,
	}
}

func (hpi highPriceIndicator) Calculate(index int) *decimal.Big {
	return new(decimal.Big).Copy(hpi.Candles[index].MaxPrice)
}

type lowPriceIndicator struct {
	*TimeSeries
}

// NewLowPriceIndicator returns an Indicator which returns the low price of a candle for a given index
func NewLowPriceIndicator(series *TimeSeries) Indicator {
	return lowPriceIndicator{
		series,
	}
}

func (lpi lowPriceIndicator) Calculate(index int) *decimal.Big {
	return new(decimal.Big).Copy(lpi.Candles[index].MinPrice)
}

type openPriceIndicator struct {
	*TimeSeries
}

// NewOpenPriceIndicator returns an Indicator which returns the open price of a candle for a given index
func NewOpenPriceIndicator(series *TimeSeries) Indicator {
	return openPriceIndicator{
		series,
	}
}

func (opi openPriceIndicator) Calculate(index int) *decimal.Big {
	return new(decimal.Big).Copy(opi.Candles[index].OpenPrice)
}

type typicalPriceIndicator struct {
	*TimeSeries
}

// NewTypicalPriceIndicator returns an Indicator which returns the typical price of a candle for a given index.
// The typical price is an average of the high, low, and close prices for a given candle.
func NewTypicalPriceIndicator(series *TimeSeries) Indicator {
	return typicalPriceIndicator{series}
}

func (tpi typicalPriceIndicator) Calculate(index int) *decimal.Big {
	tmp := new(decimal.Big).Copy(tpi.Candles[index].MaxPrice)
	tmp1 := tmp.Add(tmp, tpi.Candles[index].MinPrice)
	numerator := tmp1.Add(tmp1, tpi.Candles[index].ClosePrice)
	return numerator.Quo(numerator, decimal.New(3, 0))
}
