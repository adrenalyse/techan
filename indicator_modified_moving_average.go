package techan

import (
	"github.com/dgraph-io/ristretto"
	"github.com/sdcoffey/big"
	"log"
)

type modifiedMovingAverageIndicator struct {
	indicator   Indicator
	window      int
	resultCache *ristretto.Cache
}

// NewMMAIndicator returns a derivative indciator which returns the modified moving average of the underlying
// indictator. An in-depth explanation can be found here:
// https://en.wikipedia.org/wiki/Moving_average#Modified_moving_average
func NewMMAIndicator(indicator Indicator, window int) Indicator {
	cache := NewRistrettoCache()

	return &modifiedMovingAverageIndicator{
		indicator:   indicator,
		window:      window,
		resultCache: cache,
	}
}

func (mma *modifiedMovingAverageIndicator) Calculate(index int) big.Decimal {
	if index < mma.window-1 {
		return big.ZERO
	} else if index == mma.window-1 {
		v := NewSimpleMovingAverage(mma.indicator, mma.window).Calculate(index)
		if ok := mma.resultCache.Set(index, v, 1); !ok {
			log.Println(ok)
		}
		return v
	}
	value, ok := mma.resultCache.Get(index)
	if value != nil || ok {
		v := value.(big.Decimal)
		return v
	}

	todayVal := mma.indicator.Calculate(index)
	lastVal := mma.Calculate(index - 1)

	result := lastVal.Add(big.NewDecimal(1.0 / float64(mma.window)).Mul(todayVal.Sub(lastVal)))

	if ok = mma.resultCache.Set(index, result, 1); !ok {
		log.Println(ok)
	}

	return result
}

//func (mma modifiedMovingAverageIndicator) cache() resultCache {
//	return mma.resultCache
//}

//func (mma *modifiedMovingAverageIndicator) setCache(cache resultCache) {
//	mma.resultCache = cache
//}

//func (mma modifiedMovingAverageIndicator) windowSize() int {
//	return mma.window
//}
