package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/adshao/go-binance/v2"
)

var binance_client *binance.Client

type BinanceExchangeRepository struct{}

func (e *BinanceExchangeRepository) Init() {
	apiKey := os.Getenv("BINANCE_API_KEY")
	secretKey := os.Getenv("BINANCE_SECRET_KEY")
	binance.UseTestnet = true
	binance_client = binance.NewClient(apiKey, secretKey)
}

func (e BinanceExchangeRepository) Name() string {
	return "binanceus"
}

func (e BinanceExchangeRepository) GetCandles(ticker string, interval string, limit int) ([]Candle, error) {
	fmt.Printf("Candles for: %v, interval: %v\n\n", ticker, interval)

	klines, err := binance_client.NewKlinesService().Symbol(ticker).Interval(interval).Limit(limit).Do(context.Background())

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
	res, err := binance_client.NewExchangeInfoService().Do(context.Background())

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

func (e BinanceExchangeRepository) IsSymbolTradable(symbol string, exchange_info []ExchangeInfoSymbol) bool {
	for _, t := range exchange_info {
		if symbol != t.symbol {
			continue
		}

		return t.status == "TRADING"
	}

	return false
}
