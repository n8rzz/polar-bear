package main

import (
	"fmt"
)

type Candle struct {
	close     string
	high      string
	low       string
	open      string
	open_time int64
	volume    string
}

type ExchangeInfoSymbol struct {
	symbol      string
	status      string
	base_asset  string
	quote_asset string
}

type CandleRequest struct {
	interval string
	limit    int
}

type ExchangeRepository interface {
	// BuyLimit()
	// BuyMarket()
	// GetBalance()
	GetCandles(ticker string, interval string, limit int) ([]Candle, error)
	GetExchangeInfo() []ExchangeInfoSymbol
	// GetMarkets()
	// GetMarketSummary()
	// GetOrderBook()
	// GetTicker()
	Init()
	Name() string
	// SellLimit()
	// SellMarket()
}

func FetchCandleDataAndGenerateSignals(repository ExchangeRepository) map[string][]Candle {
	repository.Init()
	symbols := repository.GetExchangeInfo()

	fmt.Printf("Scanning %v symbols from %v exchange\n\n", len(symbols), repository.Name())

	ticker_candles := make(map[string][]Candle, len(symbols))

	for _, e := range symbols {
		// goroutine
		req := CandleRequest{interval: "15m", limit: 1000}
		candles, err := repository.GetCandles(e.symbol, req.interval, req.limit)

		if err != nil {
			fmt.Println(err)
		}

		ticker_candles[e.symbol] = candles
	}

	return ticker_candles
}
