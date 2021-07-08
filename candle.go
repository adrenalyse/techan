package techan

import (
	"fmt"
	"github.com/ericlagergren/decimal"
	"strings"
)

// Candle represents basic market information for a security over a given time period
type Candle struct {
	Period     TimePeriod
	OpenPrice  decimal.Big
	ClosePrice decimal.Big
	MaxPrice   decimal.Big
	MinPrice   decimal.Big
	Volume     decimal.Big
	TradeCount uint
}

// NewCandle returns a new *Candle for a given time period
func NewCandle(period TimePeriod) (c *Candle) {
	return &Candle{
		Period:     period,
		OpenPrice:  decimal.Big{},
		ClosePrice: decimal.Big{},
		MaxPrice:   decimal.Big{},
		MinPrice:   decimal.Big{},
		Volume:     decimal.Big{},
	}
}

// AddTrade adds a trade to this candle. It will determine if the current price is higher or lower than the min or max
// price and increment the tradecount.
func (c *Candle) AddTrade(tradeAmount, tradePrice decimal.Big) {
	if c.OpenPrice.Sign() == 0 {
		c.OpenPrice = tradePrice
	}
	c.ClosePrice = tradePrice

	if c.MaxPrice.Sign() == 0 {
		c.MaxPrice = tradePrice
	} else if tradePrice.Cmp(&c.MaxPrice) == 1 {
		c.MaxPrice = tradePrice
	}

	if c.MinPrice.Sign() == 0 {
		c.MinPrice = tradePrice
	} else if tradePrice.Cmp(&c.MinPrice) == -1 {
		c.MinPrice = tradePrice
	}

	if c.Volume.Sign() == 0 {
		c.Volume = tradeAmount
	} else {
		c.Volume = *c.Volume.Add(&c.Volume, &tradeAmount)
	}

	c.TradeCount++
}

func (c *Candle) String() string {
	return strings.TrimSpace(fmt.Sprintf(
		`
Time:	%s
Open:	%s
Close:	%s
High:	%s
Low:	%s
Volume:	%s
	`,
		c.Period,
		c.OpenPrice.String(),
		c.ClosePrice.String(),
		c.MaxPrice.String(),
		c.MinPrice.String(),
		c.Volume.String(),
	))
}
