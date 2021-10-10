package main

import "github.com/adshao/go-binance/v2"

type ExchangeRepository interface {
	// BuyLimit()
	// BuyMarket()
	// GetBalance()
	GetCandles(ticker string) []*binance.Kline
	GetExchangeInfo() (res *binance.ExchangeInfo, err error)
	// GetMarkets()
	// GetMarketSummary()
	// GetOrderBook()
	// GetTicker()
	Init()
	Name() string
	// SellLimit()
	// SellMarket()
}
