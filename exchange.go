package main

import "github.com/adshao/go-binance/v2"

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
	GetCandles(ticker string) []*binance.Kline
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
