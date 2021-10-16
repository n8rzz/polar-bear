package main

import (
	"fmt"
	"log"

	"github.com/adshao/go-binance/v2"
)

var binance_client *binance.Client

type BinanceExchangeRepository struct {
	service *BinanceService
}

func (e *BinanceExchangeRepository) Init(service *BinanceService) {
	e.service = service

	e.service.Init()
}

func (e BinanceExchangeRepository) Name() string {
	return "binanceus"
}

func (e BinanceExchangeRepository) GetCandles(ticker string, interval string, limit int) ([]Candle, error) {
	fmt.Printf("Candles for: %v, interval: %v\n\n", ticker, interval)

	klines, err := e.service.GetCandles(ticker, interval, limit)

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
	res, err := e.service.GetExchangeInfo()

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
