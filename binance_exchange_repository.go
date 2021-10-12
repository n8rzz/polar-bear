package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/adshao/go-binance/v2"
)

var client *binance.Client

type BinanceExchangeRepository struct{}

func (e *BinanceExchangeRepository) Init() {
	apiKey := os.Getenv("BINANCE_API_KEY")
	secretKey := os.Getenv("BINANCE_SECRET_KEY")
	binance.UseTestnet = true
	client = binance.NewClient(apiKey, secretKey)
}

func (e BinanceExchangeRepository) Name() string {
	return "binanceus"
}

func (e BinanceExchangeRepository) GetCandles(ticker string, interval string, limit int) ([]Candle, error) {
	fmt.Printf("Candles for: %v, interval: %v\n\n", ticker, interval)

	klines, err := client.NewKlinesService().Symbol(ticker).Interval(interval).Limit(limit).Do(context.Background())

	if err != nil {
		return nil, err
	}

	candles := make([]Candle, len(klines))

	for i, k := range klines {
		candles[i] = Candle{
			close:     k.Close,
			high:      k.High,
			low:       k.Low,
			open:      k.Open,
			open_time: k.OpenTime / 1000,
			volume:    k.Volume,
		}
	}

	return candles, nil
}

func (e BinanceExchangeRepository) GetExchangeInfo() []ExchangeInfoSymbol {
	res, err := client.NewExchangeInfoService().Do(context.Background())

	if err != nil {
		log.Fatal(err)
	}

	available_exchange_symbols := make([]ExchangeInfoSymbol, len(res.Symbols))

	for i, s := range res.Symbols {
		e := ExchangeInfoSymbol{
			s.Symbol,
			s.Status,
			s.BaseAsset,
			s.QuoteAsset,
		}
		available_exchange_symbols[i] = e
	}

	return available_exchange_symbols
}
