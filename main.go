package main

import (
	"encoding/json"
	"fmt"
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

	botConfigs := LoadBotConfig()
	strategyConfigs := LoadStrategy()
	bot := botConfigs.NewBotFromConfig(strategyConfigs)

	s, err := json.MarshalIndent(bot, "", "\t")
	fmt.Printf("Bot: %+v\n\n", string(s))
	fmt.Print("\n\n============\n")

	if err != nil {
		log.Fatal(err)
	}

	repository := &BinanceExchangeRepository{}
	service := &BinanceService{}

	repository.Init(service)

	ticker_candles := FetchCandleDataAndGenerateSignals(bot, repository)

	for k := range ticker_candles {
		fmt.Print("\n\n------------\n\n")
		fmt.Printf("CalculateSignals for %s\n", k)
		CalculateSignals(ticker_candles[k], bot)
		fmt.Print("\n+++ +++ +++\n\n")
	}
}
