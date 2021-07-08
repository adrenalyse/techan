package techan

import (
	"github.com/ericlagergren/decimal"
)

type modifiedMovingAverageIndicator struct {
	indicator   Indicator
	window      int
	resultCache resultCache
}

// NewMMAIndicator returns a derivative indicator which returns the modified moving average of the underlying
// indicator. An in-depth explanation can be found here:
// https://en.wikipedia.org/wiki/Moving_average#Modified_moving_average
func NewMMAIndicator(indicator Indicator, window int) Indicator {
	return &modifiedMovingAverageIndicator{
		indicator:   indicator,
		window:      window,
		resultCache: make([]*decimal.Big, 10000),
	}
}

func (mma *modifiedMovingAverageIndicator) Calculate(index int) decimal.Big {
	if cachedValue := returnIfCached(mma, index, func(i int) decimal.Big {
		return NewSimpleMovingAverage(mma.indicator, mma.window).Calculate(i)
	}); cachedValue != nil {
		return *cachedValue
	}

	todayVal := mma.indicator.Calculate(index)
	lastVal := mma.Calculate(index - 1)

	// lastVal + (big.NewDecimal(1.0 / float64(mma.window)) * todayVal.Sub(lastVal)
	tmp := new(decimal.Big).SetFloat64(1.0 / float64(mma.window))
	r := new(decimal.Big).FMA(tmp, todayVal.Sub(&todayVal, &lastVal), &lastVal)
	cacheResult(mma, index, *r)

	return *r
}

func (mma modifiedMovingAverageIndicator) cache() resultCache {
	return mma.resultCache
}

func (mma *modifiedMovingAverageIndicator) setCache(cache resultCache) {
	mma.resultCache = cache
}

func (mma modifiedMovingAverageIndicator) windowSize() int {
	return mma.window
}
