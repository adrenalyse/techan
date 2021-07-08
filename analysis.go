package techan

//lint:file-ignore S1038 prefer Fprintln

import (
	"fmt"
	"github.com/ericlagergren/decimal"
	"io"
	"time"
)

// Analysis is an interface that describes a methodology for taking a TradingRecord as input,
// and giving back some float value that describes it's performance with respect to that methodology.
type Analysis interface {
	Analyze(*TradingRecord) float64
}

// TotalProfitAnalysis analyzes the trading record for total profit.
type TotalProfitAnalysis struct{}

// Analyze analyzes the trading record for total profit.
func (tps TotalProfitAnalysis) Analyze(record *TradingRecord) float64 {
	totalProfit := &decimal.Big{}
	for _, trade := range record.Trades {
		if trade.IsClosed() {

			costBasis := trade.CostBasis()
			exitValue := trade.ExitValue()

			if trade.IsLong() {
				totalProfit.Add(totalProfit, exitValue.Sub(exitValue, costBasis))
			} else if trade.IsShort() {
				totalProfit.Sub(totalProfit, exitValue.Sub(exitValue, costBasis))
			}

		}
	}

	f, _ := totalProfit.Float64()
	return f
}

// PercentGainAnalysis analyzes the trading record for the percentage profit gained relative to start
type PercentGainAnalysis struct{}

// Analyze analyzes the trading record for the percentage profit gained relative to start
func (pga PercentGainAnalysis) Analyze(record *TradingRecord) float64 {
	if len(record.Trades) > 0 && record.Trades[0].IsClosed() {
		tmp := record.Trades[len(record.Trades)-1].ExitValue()
		tmp1 := tmp.Quo(tmp, record.Trades[0].CostBasis())
		f, _ := tmp1.Sub(tmp1, decimal.New(1, 0)).Float64()
		return f
	}

	return 0
}

// NumTradesAnalysis analyzes the trading record for the number of trades executed
type NumTradesAnalysis string

// Analyze analyzes the trading record for the number of trades executed
func (nta NumTradesAnalysis) Analyze(record *TradingRecord) float64 {
	return float64(len(record.Trades))
}

// LogTradesAnalysis is a wrapper around an io.Writer, which logs every trade executed to that writer
type LogTradesAnalysis struct {
	io.Writer
}

// Analyze logs trades to provided io.Writer
func (lta LogTradesAnalysis) Analyze(record *TradingRecord) float64 {
	logOrder := func(trade *Position) {
		fmt.Fprintln(lta.Writer, fmt.Sprintf("%s - enter with buy %s (%s @ $%s)", trade.EntranceOrder().ExecutionTime.UTC().Format(time.RFC822), trade.EntranceOrder().Security, trade.EntranceOrder().Amount, trade.EntranceOrder().Price))
		fmt.Fprintln(lta.Writer, fmt.Sprintf("%s - exit with sell %s (%s @ $%s)", trade.ExitOrder().ExecutionTime.UTC().Format(time.RFC822), trade.ExitOrder().Security, trade.ExitOrder().Amount, trade.ExitOrder().Price))

		tmp := trade.ExitValue()
		profit := tmp.Sub(tmp, trade.CostBasis())
		fmt.Fprintln(lta.Writer, fmt.Sprintf("Profit: $%s", profit))
	}

	for _, trade := range record.Trades {
		if trade.IsClosed() {
			logOrder(trade)
		}
	}
	return 0.0
}

// PeriodProfitAnalysis analyzes the trading record for the average profit based on the time period provided.
// i.e., if the trading record spans a year of trading, and PeriodProfitAnalysis wraps one month, Analyze will return
// the total profit for the whole time period divided by 12.
type PeriodProfitAnalysis struct {
	Period time.Duration
}

// Analyze returns the average profit for the trading record based on the given duration
func (ppa PeriodProfitAnalysis) Analyze(record *TradingRecord) float64 {
	var tp TotalProfitAnalysis
	totalProfit := tp.Analyze(record)

	periods := record.Trades[len(record.Trades)-1].ExitOrder().ExecutionTime.Sub(record.Trades[0].EntranceOrder().ExecutionTime) / ppa.Period
	return totalProfit / float64(periods)
}

// ProfitableTradesAnalysis analyzes the trading record for the number of profitable trades
type ProfitableTradesAnalysis struct{}

// Analyze returns the number of profitable trades in a trading record
func (pta ProfitableTradesAnalysis) Analyze(record *TradingRecord) float64 {
	var profitableTrades int
	tmp := new(decimal.Big)
	tmp1 := new(decimal.Big)
	for _, trade := range record.Trades {
		tmp.Set(trade.EntranceOrder().Amount)
		costBasis := tmp.Mul(tmp, trade.EntranceOrder().Price)
		tmp1.Set(trade.ExitOrder().Amount)
		sellPrice := tmp1.Mul(tmp1, trade.ExitOrder().Price)

		if sellPrice.Cmp(costBasis) == 1 {
			profitableTrades++
		}
	}

	return float64(profitableTrades)
}

// AverageProfitAnalysis returns the average profit for the trading record. Average profit is represented as the total
// profit divided by the number of trades executed.
type AverageProfitAnalysis struct{}

// Analyze returns the average profit of the trading record
func (apa AverageProfitAnalysis) Analyze(record *TradingRecord) float64 {
	var tp TotalProfitAnalysis
	totalProft := tp.Analyze(record)

	return totalProft / float64(len(record.Trades))
}

// BuyAndHoldAnalysis returns the profit based on a hypothetical where a purchase order was made on the first period available
// and held until the date on the last trade of the trading record. It's useful for comparing the performance of your strategy
// against a simple long position.
type BuyAndHoldAnalysis struct {
	TimeSeries    *TimeSeries
	StartingMoney float64
}

// Analyze returns the profit based on a simple buy and hold strategy
func (baha BuyAndHoldAnalysis) Analyze(record *TradingRecord) float64 {
	if len(record.Trades) == 0 {
		return 0
	}

	tmp := new(decimal.Big).SetFloat64(baha.StartingMoney)
	openOrder := Order{
		Side:   BUY,
		Amount: tmp.Quo(tmp, baha.TimeSeries.Candles[0].ClosePrice),
		Price:  baha.TimeSeries.Candles[0].ClosePrice,
	}

	closeOrder := Order{
		Side:   SELL,
		Amount: openOrder.Amount,
		Price:  baha.TimeSeries.Candles[len(baha.TimeSeries.Candles)-1].ClosePrice,
	}

	pos := NewPosition(openOrder)
	pos.Exit(closeOrder)

	tmp = pos.ExitValue()
	f, _ := tmp.Sub(tmp, pos.CostBasis()).Float64()
	return f
}
