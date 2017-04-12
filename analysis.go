package talib4g

import (
	"fmt"
	"log"
	"time"
)

type Analysis interface {
	Analyze(*TradingRecord) float64
}

type TotalProfitAnalysis string

func (tps TotalProfitAnalysis) Analyze(record *TradingRecord) float64 {
	profit := 0.0
	for _, trade := range record.Trades {
		costBasis := trade.EntranceOrder().Amount * trade.EntranceOrder().Price
		sellPrice := trade.ExitOrder().Amount * trade.ExitOrder().Price

		profit += sellPrice - costBasis
	}

	return profit
}

type NumTradesAnalysis string

func (nta NumTradesAnalysis) Analyze(record *TradingRecord) float64 {
	return float64(len(record.Trades))
}

type LogTradesAnalysis string

func (lta LogTradesAnalysis) Analyze(record *TradingRecord) float64 {
	logOrder := func(order *Order) {
		var oType string
		var action string
		if order.Type == BUY {
			oType = "buy"
			action = "Entered"
		} else {
			oType = "sell"
			action = "Exited"
		}

		log.Println(fmt.Sprintf("%s - %s with %s (%f @ $%f)", order.ExecutionTime.Format(time.RFC822), action, oType, order.Amount, order.Price))
	}

	for _, trade := range record.Trades {
		logOrder(trade.EntranceOrder())
		logOrder(trade.ExitOrder())
	}
	return 0.0
}

type ProfitableTradesAnalysis string

func (pta ProfitableTradesAnalysis) Analyze(record *TradingRecord) float64 {
	var profitableTrades int
	for _, trade := range record.Trades {
		costBasis := trade.EntranceOrder().Amount * trade.EntranceOrder().Price
		sellPrice := trade.ExitOrder().Amount * trade.ExitOrder().Price

		if sellPrice > costBasis {
			profitableTrades++
		}
	}

	return float64(profitableTrades)
}

type AverageProfitAnalysis string

func (apa AverageProfitAnalysis) Analyze(record *TradingRecord) float64 {
	var tp TotalProfitAnalysis
	totalProft := tp.Analyze(record)

	return totalProft / float64(len(record.Trades))
}
