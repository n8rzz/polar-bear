package main

import (
	"fmt"
	"time"

	"github.com/adshao/go-binance/v2"
	"github.com/sdcoffey/big"
	"github.com/sdcoffey/techan"
)

func CalculateSignals(candles []*binance.Kline) {
	series := techan.NewTimeSeries()

	// fmt.Printf("entry: %+v\n\n", candles[0])
	// fmt.Printf("entry: %+v\n\n", candles[len(candles)-1])

	for _, entry := range candles {
		start := time.Unix((entry.OpenTime / 1000), 0)
		end := time.Unix((entry.CloseTime / 1000), 0)
		period := techan.NewTimePeriod(start, end.Sub(start))

		candle := techan.NewCandle(period)
		candle.OpenPrice = big.NewFromString(entry.Open)
		candle.ClosePrice = big.NewFromString(entry.Close)
		candle.MaxPrice = big.NewFromString(entry.High)
		candle.MinPrice = big.NewFromString(entry.Low)

		series.AddCandle(candle)
	}

	closePrices := techan.NewClosePriceIndicator(series)
	ma := techan.NewEMAIndicator(closePrices, 9)

	fmt.Println(ma.Calculate(0).FormattedString(8))
}
