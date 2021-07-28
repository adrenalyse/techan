package techan

import (
	"github.com/dgraph-io/ristretto"
	"log"

	"github.com/sdcoffey/big"
)

type emaIndicator struct {
	indicator   Indicator
	window      int
	alpha       big.Decimal
	resultCache *ristretto.Cache
}

// NewEMAIndicator returns a derivative indicator which returns the average of the current and preceding values in
// the given windowSize, with values closer to current index given more weight. A more in-depth explanation can be found here:
// http://www.investopedia.com/terms/e/ema.asp
func NewEMAIndicator(indicator Indicator, window int) Indicator {
	cache := NewRistrettoCache()

	return &emaIndicator{
		indicator:   indicator,
		window:      window,
		alpha:       big.ONE.Frac(2).Div(big.NewFromInt(window + 1)),
		resultCache: cache,
	}
}

func (ema *emaIndicator) Calculate(index int) big.Decimal {
	if index < ema.window-1 {
		return big.ZERO
	} else if index == ema.window-1 {
		v := NewSimpleMovingAverage(ema.indicator, ema.window).Calculate(index)
		if ok := ema.resultCache.Set(index, v, 1); !ok {
			log.Println(ok)
		}
		return v
	}
	value, ok := ema.resultCache.Get(index)
	if value != nil || ok {
		v := value.(big.Decimal)
		return v
	}
	//if cachedValue := returnIfCached(ema, index, func(i int) big.Decimal {
	//	return NewSimpleMovingAverage(ema.indicator, ema.window).Calculate(i)
	//}); cachedValue != nil {
	//	return *cachedValue
	//}

	todayVal := ema.indicator.Calculate(index).Mul(ema.alpha)
	result := todayVal.Add(ema.Calculate(index - 1).Mul(big.ONE.Sub(ema.alpha)))

	if ok = ema.resultCache.Set(index, result, 1); !ok {
		log.Println(ok)
	}

	return result
}

//func (ema emaIndicator) cache() resultCache { return ema.resultCache }

//func (ema *emaIndicator) setCache(newCache resultCache) {
//	ema.resultCache = newCache
//}

//func (ema emaIndicator) windowSize() int { return ema.window }
