package techan

import (
	"github.com/ericlagergren/decimal"
	"time"
)

// OrderSide is a simple enumeration representing the side of an Order (buy or sell)
type OrderSide int

// BUY and SELL enumerations
const (
	BUY OrderSide = iota
	SELL
)

// Order represents a trade execution (buy or sell) with associated metadata.
type Order struct {
	Side          OrderSide
	Security      string
	Price         *decimal.Big
	Amount        *decimal.Big
	ExecutionTime time.Time
}
