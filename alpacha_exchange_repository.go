package main

import (
	"fmt"
	"log"
	"os"

	"github.com/alpacahq/alpaca-trade-api-go/v2/alpaca"
)

type AlpachaExchangeRepository struct{}

func (e AlpachaExchangeRepository) GetCandles(ticker string, interval string, limit int) ([]Candle, error) {
	return []Candle{}, nil
}

func (e AlpachaExchangeRepository) GetExchangeInfo() []ExchangeInfoSymbol {
	return []ExchangeInfoSymbol{}
}

func (e AlpachaExchangeRepository) Init() {
	apiKey := os.Getenv("ALPACHA_API_KEY")
	secretKey := os.Getenv("ALPACHA_SECRET_KEY")
	client := alpaca.NewClient(alpaca.ClientOpts{
		ApiKey:    apiKey,
		ApiSecret: secretKey,
		BaseURL:   "https://paper-api.alpaca.markets",
	})

	acct, err := client.GetAccount()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", *acct)
}

func (e AlpachaExchangeRepository) Name() string {
	return "alpacha"
}
