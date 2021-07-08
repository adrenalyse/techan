package techan

import (
	"fmt"
	"github.com/ericlagergren/decimal"
	"math/rand"
	"testing"
	"time"

	"strconv"

	"github.com/stretchr/testify/assert"
)

var candleIndex int
var mockedTimeSeries = mockTimeSeriesFl(
	64.75, 63.79, 63.73,
	63.73, 63.55, 63.19,
	63.91, 63.85, 62.95,
	63.37, 61.33, 61.51)

func randomTimeSeries(size int) *TimeSeries {
	vals := make([]string, size)
	rand.Seed(time.Now().Unix())
	for i := 0; i < size; i++ {
		val := rand.Float64() * 100
		if i == 0 {
			vals[i] = fmt.Sprint(val)
		} else {
			last, _ := strconv.ParseFloat(vals[i-1], 64)
			if i%2 == 0 {
				vals[i] = fmt.Sprint(last + (val / 10))
			} else {
				vals[i] = fmt.Sprint(last - (val / 10))
			}
		}
	}

	return mockTimeSeries(vals...)
}

func mockTimeSeriesOCHL(values ...[]float64) *TimeSeries {
	ts := NewTimeSeries()
	for i, ochl := range values {
		candle := NewCandle(NewTimePeriod(time.Unix(int64(i), 0), time.Second))
		candle.OpenPrice = *new(decimal.Big).SetFloat64(ochl[0])
		candle.ClosePrice = *new(decimal.Big).SetFloat64(ochl[1])
		candle.MaxPrice = *new(decimal.Big).SetFloat64(ochl[2])
		candle.MinPrice = *new(decimal.Big).SetFloat64(ochl[3])
		candle.Volume = *new(decimal.Big).SetFloat64(float64(i))

		ts.AddCandle(candle)
	}

	return ts
}

func mockTimeSeries(values ...string) *TimeSeries {
	ts := NewTimeSeries()
	for _, val := range values {
		candle := NewCandle(NewTimePeriod(time.Unix(int64(candleIndex), 0), time.Second))
		op, _ := new(decimal.Big).SetString(val)
		candle.OpenPrice = *op
		cp, _ := new(decimal.Big).SetString(val)
		candle.ClosePrice = *cp
		tmp, _ := new(decimal.Big).SetString(val)
		candle.MaxPrice = *tmp.Add(tmp, decimal.New(1, 0))
		tmp1, _ := new(decimal.Big).SetString(val)
		candle.MinPrice = *tmp1.Sub(tmp1, decimal.New(1, 0))
		v, _ := new(decimal.Big).SetString(val)
		candle.Volume = *v

		ts.AddCandle(candle)

		candleIndex++
	}

	return ts
}

func mockTimeSeriesFl(values ...float64) *TimeSeries {
	strVals := make([]string, len(values))

	for i, val := range values {
		strVals[i] = fmt.Sprint(val)
	}

	return mockTimeSeries(strVals...)
}

func decimalEquals(t *testing.T, expected float64, actual decimal.Big) {
	f, _ := actual.Float64()
	assert.Equal(t, fmt.Sprintf("%.4f", expected), fmt.Sprintf("%.4f", f))
}

func dump(indicator Indicator) (values []float64) {

	defer func() {
		recover()
	}()

	var index int
	for {
		c := indicator.Calculate(index)
		f, _ := c.Float64()
		//log.Println(f)
		values = append(values, f)
		index++
	}

	return
}

func indicatorEquals(t *testing.T, expected []float64, indicator Indicator) {
	actualValues := dump(indicator)
	assert.EqualValues(t, expected, actualValues)
}
