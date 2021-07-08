package techan

import (
	"github.com/ericlagergren/decimal"
)

type emaIndicator struct {
	indicator   Indicator
	window      int
	alpha       decimal.Big
	resultCache resultCache
}

// NewEMAIndicator returns a derivative indicator which returns the average of the current and preceding values in
// the given windowSize, with values closer to current index given more weight. A more in-depth explanation can be found here:
// http://www.investopedia.com/terms/e/ema.asp
func NewEMAIndicator(indicator Indicator, window int) Indicator {
	tmp := decimal.New(2, 0)
	return &emaIndicator{
		indicator:   indicator,
		window:      window,
		alpha:       *tmp.Quo(tmp, new(decimal.Big).SetUint64(uint64(window+1))),
		resultCache: make([]*decimal.Big, 1000),
	}
}

func (ema *emaIndicator) Calculate(index int) decimal.Big {
	if cachedValue := returnIfCached(ema, index, func(i int) decimal.Big {
		return NewSimpleMovingAverage(ema.indicator, ema.window).Calculate(i)
	}); cachedValue != nil {
		return *cachedValue
	}

	tmp := ema.indicator.Calculate(index)
	todayVal := tmp.Mul(&tmp, &ema.alpha)
	tmp1 := ema.Calculate(index - 1)
	result := new(decimal.Big).Add(todayVal, tmp1.Mul(&tmp1, &ema.alpha))

	cacheResult(ema, index, *result)

	return *result
}

func (ema emaIndicator) cache() resultCache { return ema.resultCache }

func (ema *emaIndicator) setCache(newCache resultCache) {
	ema.resultCache = newCache
}

func (ema emaIndicator) windowSize() int { return ema.window }
