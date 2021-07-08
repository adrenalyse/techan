package techan

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCommidityChannelIndexIndicator_Calculate(t *testing.T) {
	typicalPrices := []string{
		"23.98", "23.92", "23.79", "23.67", "23.54",
		"23.36", "23.65", "23.72", "24.16", "23.91",
		"23.81", "23.92", "23.74", "24.68", "24.94",
		"24.93", "25.10", "25.12", "25.20", "25.06",
		"24.50", "24.31", "24.57", "24.62", "24.49",
		"24.37", "24.41", "24.35", "23.75", "24.09",
	}

	series := mockTimeSeries(typicalPrices...)

	cci := NewCCIIndicator(series, 20)

	results := []string{"101.9184652278177", "31.19461183977313", "6.557771560930121", "33.60780725831046", "34.96855345911950", "13.60267993097147",
		"-10.67887109077040", "-11.47098515519568", "-29.25666070527812", "-128.6000174018968", "-72.72727272727273"}

	for i, result := range results {
		assert.EqualValues(t, result, cci.Calculate(i+19).String())
	}
}
