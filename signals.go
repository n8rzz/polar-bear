package main

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/markcheno/go-talib"
)

type SignalResultEnum string

const (
	Buy  SignalResultEnum = "Buy"
	Sell SignalResultEnum = "Sell"
)

type SignalParams struct {
	AnalyisTime   int
	SegmentParams StrategySegmentParams
}

type Signal struct {
	Exchange   []string
	Indicator  string
	IsRequired bool
	Params     SignalParams
	Result     SignalResultEnum
	Ticker     string
}

func IsFastAboveSlowForLength(slow, fast []float64, length int) bool {
	slow_index := len(slow) - 1
	fast_index := len(fast) - 1

	for i := 0; i < length; i++ {
		slow_value_to_check := []float64{slow[slow_index-i]}
		fast_value_to_check := []float64{fast[fast_index-i]}
		result := IsFastAboveSlow(slow_value_to_check, fast_value_to_check)

		if !result {
			return false
		}
	}

	return true
}

func IsFastAboveSlow(slow, fast []float64) bool {
	slow_length := len(slow)
	fast_length := len(fast)

	return fast[fast_length-1] > slow[slow_length-1]
}

func ApoSingalProcessor(segment StrategySegment, ticker string, close_prices []float64, bot *Bot) Signal {
	apo := talib.Apo(close_prices, segment.Params.FastPeriod, segment.Params.SlowPeriod, talib.EMA)
	result := IsFastAboveSlowForLength([]float64{0}, apo, 1)
	resultAction := Buy

	if !result {
		resultAction = Sell
	}

	signalParams := SignalParams{
		AnalyisTime:   time.Now().Local().UTC().Second(),
		SegmentParams: segment.Params,
	}
	signal := Signal{
		Exchange:   bot.Config.Exchange,
		Indicator:  "apo",
		IsRequired: segment.IsRequired,
		Params:     signalParams,
		Result:     resultAction,
		Ticker:     ticker,
	}

	// fmt.Printf("apo: %f\n", apo[len(apo)-5:len(apo)-1])
	// fmt.Printf("APO - >0: %v\n", IsFastAboveSlowForLength([]float64{0}, apo, 1))
	fmt.Printf("Signal::apo %v\n", signal)

	return signal
}

func RsiSingalProcessor(segment StrategySegment, ticker string, close_prices []float64, bot *Bot) Signal {
	rsi := talib.Rsi(close_prices, segment.Params.Period)
	rsi_slow_value_comparison := make([]float64, segment.KeepSignalLength)

	for i := range rsi_slow_value_comparison {
		rsi_slow_value_comparison[i] = segment.Params.SignalWhen.Value
	}

	result := IsFastAboveSlowForLength(rsi_slow_value_comparison, rsi, segment.KeepSignalLength)
	resultAction := Buy

	if !result {
		resultAction = Sell
	}

	signalParams := SignalParams{
		AnalyisTime:   time.Now().Local().UTC().Second(),
		SegmentParams: segment.Params,
	}
	signal := Signal{
		Exchange:   bot.Config.Exchange,
		Indicator:  "rsi",
		IsRequired: segment.IsRequired,
		Params:     signalParams,
		Result:     resultAction,
		Ticker:     ticker,
	}

	fmt.Printf("rsi: %f\n", rsi[len(rsi)-5:len(rsi)-1])
	// fmt.Printf("RSI - >70: %v\n", IsFastAboveSlowForLength([]float64{70}, rsi, 1))

	fmt.Printf("Signal::rsi %v\n", signal)

	return signal
}

func SingalProcessorFactory(ticker string, close_prices []float64, bot *Bot) {
	signals := make([]Signal, len(bot.Strategy.Segments))

	for i, segment := range bot.Strategy.Segments {
		fmt.Printf("\n\n")
		switch segment.Indicator {
		case "apo":
			signals[i] = ApoSingalProcessor(segment, ticker, close_prices, bot)
		case "ema":
			ema_fast := talib.Ema(close_prices, segment.Params.FastPeriod)
			ema_slow := talib.Ema(close_prices, segment.Params.SlowPeriod)

			fmt.Printf("ema_slow: %f\n", ema_slow[len(ema_slow)-5:len(ema_slow)-1])
			fmt.Printf("ema_fast: %f\n", ema_fast[len(ema_fast)-5:len(ema_fast)-1])
			// fmt.Printf("EMA - IsFastAboveSlow: %v\n", IsFastAboveSlowForLength(ema_slow, ema_fast, 1))
		case "rsi":
			signals[i] = RsiSingalProcessor(segment, ticker, close_prices, bot)
		case "sma":
			sma_fast := talib.Sma(close_prices, segment.Params.FastPeriod)
			sma_slow := talib.Sma(close_prices, segment.Params.SlowPeriod)

			fmt.Printf("sma_slow: %f\n", sma_slow[len(sma_slow)-5:len(sma_slow)-1])
			fmt.Printf("sma_fast: %f\n", sma_fast[len(sma_fast)-5:len(sma_fast)-1])
		}
	}
}

func CalculateSignals(ticker string, candles []Candle, bot *Bot) {
	close_prices := make([]float64, len(candles))

	for i, entry := range candles {
		// start := time.Unix((entry.OpenTime / 1000), 0)
		// end := time.Unix((entry.CloseTime / 1000), 0)
		close_as_float, err := strconv.ParseFloat(entry.close, 64)

		if err != nil {
			log.Fatal(err)
		}

		close_prices[i] = close_as_float
	}

	SingalProcessorFactory(ticker, close_prices, bot)
}
