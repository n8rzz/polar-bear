package main

import (
	"context"
	"fmt"
	"os"

	"github.com/adshao/go-binance/v2"
)

type ExchangeSymbol struct {
	symbol      string
	status      string
	base_asset  string
	quote_asset string
}

type CandleRequest struct {
	ticker   string
	interval string
	limit    int
}

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
	klines, err := e.client.NewKlinesService().Symbol(ticker).Interval(interval).Limit(limit).Do(context.Background())

	return klines, err
}

func (e BinanceExchangeRepository) GetExchangeInfo() (*binance.ExchangeInfo, error) {
	res, err := e.client.NewExchangeInfoService().Do(context.Background())

	if err != nil {
		return nil, err
	}

	available_exchange_symbols := make([]ExchangeSymbol, len(res.Symbols))

	for i, s := range res.Symbols {
		e := ExchangeSymbol{
			s.Symbol,
			s.Status,
			s.BaseAsset,
			s.QuoteAsset,
		}
		available_exchange_symbols[i] = e
	}

	fmt.Printf("available_exchange_symbols: %+v\n\n", available_exchange_symbols)

	return res, nil
}

func FetchCandleDataAndGenerateSignals(req CandleRequest, repository *BinanceExchangeRepository) {
	repository.Init()
	// repository.GetExchangeInfo()

	candles, err := repository.GetCandles(req.ticker, req.interval, req.limit)

	if err != nil {
		fmt.Println(err)
	}

	CalculateSignals(candles)
}
