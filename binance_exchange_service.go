package main

import (
	"context"
	"fmt"
	"os"

	"github.com/adshao/go-binance/v2"
)

type BinanceExchangeService interface {
	Init()
	GetCandles(ticker string, interval string, limit int) ([]*binance.Kline, error)
	GetExchangeInfo() (*binance.ExchangeInfo, error)
}

type BinanceService struct {
	client *binance.Client
}

func (s *BinanceService) Init() {
	apiKey := os.Getenv("BINANCE_API_KEY")
	secretKey := os.Getenv("BINANCE_SECRET_KEY")
	binance.UseTestnet = true

	s.client = binance.NewClient(apiKey, secretKey)
}

func (s *BinanceService) GetCandles(ticker string, interval string, limit int) ([]*binance.Kline, error) {
	klines, err := s.client.NewKlinesService().Symbol(ticker).Interval(interval).Limit(limit).Do(context.Background())

	if err != nil {
		fmt.Printf("error fetching candles: %+v", err)
	}

	return klines, err
}

func (s *BinanceService) GetExchangeInfo() (*binance.ExchangeInfo, error) {
	res, err := s.client.NewExchangeInfoService().Do(context.Background())

	if err != nil {
		fmt.Printf("error fetching exchange info: %+v", err)
	}

	return res, err
}
