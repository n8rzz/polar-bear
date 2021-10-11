package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/adshao/go-binance/v2"
)

type BinanceExchangeRepository struct {
	client *binance.Client
}

func (e *BinanceExchangeRepository) Init() {
	apiKey := os.Getenv("BINANCE_API_KEY")
	secretKey := os.Getenv("BINANCE_SECRET_KEY")
	binance.UseTestnet = true
	e.client = binance.NewClient(apiKey, secretKey)
}

func (e BinanceExchangeRepository) Name() string {
	return "binanceus"
}

func (e BinanceExchangeRepository) GetCandles(ticker string, interval string, limit int) ([]*binance.Kline, error) {
	fmt.Printf("Candles for: %v, interval: %v\n\n", ticker, interval)

	klines, err := e.client.NewKlinesService().Symbol(ticker).Interval(interval).Limit(limit).Do(context.Background())

	return klines, err
}

func (e BinanceExchangeRepository) GetExchangeInfo() []ExchangeInfoSymbol {
	res, err := e.client.NewExchangeInfoService().Do(context.Background())

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

func FetchCandleDataAndGenerateSignals(repository *BinanceExchangeRepository) {
	repository.Init()
	symbols := repository.GetExchangeInfo()

	fmt.Printf("Scanning %v symbols from %v exchange\n\n", len(symbols), repository.Name())

	for _, e := range symbols {
		// goroutine
		fmt.Println("")
		fmt.Println("++++++++++")
		req := CandleRequest{interval: "15m", limit: 1000}
		candles, err := repository.GetCandles(e.symbol, req.interval, req.limit)

		if err != nil {
			fmt.Println(err)
		}

		CalculateSignals(candles)

		fmt.Println("")
		fmt.Println("----------")

		time.Sleep(2 * time.Second)
	}
}
