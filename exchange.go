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
	IsSymbolTradable(symbol string, exchange_info []ExchangeInfoSymbol) bool
	Init(service *BinanceService)
	Name() string
	// SellLimit()
	// SellMarket()
}

func FetchCandleDataAndGenerateSignals(bot *Bot, repository ExchangeRepository) map[string][]Candle {
	exchange_symbols := repository.GetExchangeInfo()

	fmt.Printf("exchange_symbols - %+v\n\n", exchange_symbols)
	fmt.Printf("Scanning %v symbols from %v exchange\n\n", len(exchange_symbols), repository.Name())

	ticker_candles := make(map[string][]Candle, len(bot.Config.Tickers))

	for _, t := range bot.Config.Tickers {
		// goroutine
		if repository.IsSymbolTradable(t, exchange_symbols) {
			req := CandleRequest{interval: "15m", limit: 1000}
			candles, err := repository.GetCandles(t, req.interval, req.limit)

			if err != nil {
				fmt.Println(err)
			}

			ticker_candles[t] = candles
		}
	}

	return ticker_candles
}
