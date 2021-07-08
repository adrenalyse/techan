package techan

import (
	"github.com/ericlagergren/decimal"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRelativeStrengthIndexIndicator(t *testing.T) {
	indicator := NewRelativeStrengthIndexIndicator(NewClosePriceIndicator(mockedTimeSeries), 3)

	expectedValues := []float64{
		0,
		0,
		0,
		0,
		0,
		0,
		57.99522673031028,
		54.07510431154382,
		21.45103448275862,
		44.77394829224772,
		14.15424044734388,
		21.2794458508391,
	}

	indicatorEquals(t, expectedValues, indicator)
}

func TestRelativeStrengthIndicator(t *testing.T) {
	indicator := NewRelativeStrengthIndicator(NewClosePriceIndicator(mockedTimeSeries), 3)

	expectedValues := []float64{
		0,
		0,
		0,
		0,
		0,
		0,
		1.380681818181819,
		1.177468201090249,
		0.2730912411322611,
		0.8107396221114,
		0.1648799022933912,
		0.2703162608652177,
	}

	indicatorEquals(t, expectedValues, indicator)
}

func TestRelativeStrengthIndicatorNoPriceChange(t *testing.T) {
	newClosePriceIndicator := NewClosePriceIndicator(mockTimeSeries("42.0", "42.0"))
	rsInd := NewRelativeStrengthIndicator(newClosePriceIndicator, 2)
	assert.Equal(t, *new(decimal.Big).SetInf(false), rsInd.Calculate(1))
}
