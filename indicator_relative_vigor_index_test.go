package techan

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRelativeVigorIndexIndicator_Calculate(t *testing.T) {
	series := mockTimeSeriesOCHL(
		[]float64{10, 12, 12, 8},
		[]float64{11, 14, 14, 9},
		[]float64{8, 19, 20, 8},
		[]float64{9, 10, 11, 8},
	)

	rvii := NewRelativeVigorIndexIndicator(series)

	t.Run("Returns zero when index < 4", func(t *testing.T) {
		c := rvii.Calculate(0)
		assert.EqualValues(t, "0", c.String())
		c = rvii.Calculate(1)
		assert.EqualValues(t, "0", c.String())
		c = rvii.Calculate(2)
		assert.EqualValues(t, "0", c.String())
	})

	t.Run("Calculates rvii", func(t *testing.T) {
		c := rvii.Calculate(3)
		assert.EqualValues(t, "0.7560975609756098", c.String())
	})
}

func TestRelativeVigorIndexSignalLine_Calculate(t *testing.T) {
	series := mockTimeSeriesOCHL(
		[]float64{10, 12, 12, 8},
		[]float64{11, 14, 14, 9},
		[]float64{8, 19, 20, 8},
		[]float64{9, 10, 11, 8},
		[]float64{11, 14, 14, 9},
		[]float64{9, 10, 11, 8},
		[]float64{10, 12, 12, 8},
		[]float64{9, 10, 11, 8},
	)

	signalLine := NewRelativeVigorSignalLine(series)

	t.Run("Returns zero when index < 0", func(t *testing.T) {
		c := signalLine.Calculate(0)
		assert.EqualValues(t, "0", c.String())
		c = signalLine.Calculate(1)
		assert.EqualValues(t, "0", c.String())
		c = signalLine.Calculate(2)
		assert.EqualValues(t, "0", c.String())
	})

	t.Run("Calculates rvii signal line", func(t *testing.T) {
		c := signalLine.Calculate(7)
		assert.EqualValues(t, "0.5752316290535085", c.String())
	})
}
