package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/markcheno/go-talib"
)

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

func CalculateSignals(candles []Candle, bot *Bot) {
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

	ema_fast := talib.Ema(close_prices, 9)
	ema_slow := talib.Ema(close_prices, 20)
	rsi := talib.Rsi(close_prices, 9)
	apo := talib.Apo(close_prices, 12, 26, talib.EMA)

	fmt.Printf("ema_slow: %f\n", ema_slow[len(ema_slow)-5:len(ema_slow)-1])
	fmt.Printf("ema_fast: %f\n", ema_fast[len(ema_fast)-5:len(ema_fast)-1])
	fmt.Printf("rsi: %f\n", rsi[len(rsi)-5:len(rsi)-1])
	fmt.Printf("apo: %f\n", apo[len(apo)-5:len(apo)-1])
	fmt.Print("\n\n")
	fmt.Printf("EMA - IsFastAboveSlow: %v\n", IsFastAboveSlowForLength(ema_slow, ema_fast, 1))
	fmt.Printf("RSI - >70: %v\n", IsFastAboveSlowForLength([]float64{70}, rsi, 1))
	fmt.Printf("APO - >0: %v\n", IsFastAboveSlowForLength([]float64{0}, apo, 1))
}
