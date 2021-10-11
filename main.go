package main

import (
	"log"

	"github.com/joho/godotenv"
)

/*
 this bot's work can be broken down into three phases:
 scan -> signal -> order

 more specifically:

 - collect tickers we want to monitor
 - periodically pull candles
 - check for signals for each ticker
 - when found, attempt to open position
 - monitor open order until accept
 - monitor position for closing singal
 - when found, attempt to close position
 - monitor close order until accept
*/
func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	repository := &BinanceExchangeRepository{}

	FetchCandleDataAndGenerateSignals(repository)
}
