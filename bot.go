package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

type BotConfig struct {
	Exchange             []string `json:"exchange"`
	Id                   string   `json:"id"`
	IsActive             bool     `json:"isActive"`
	IsPaper              bool     `json:"isPaper"`
	MaxAllocationAmount  float64  `json:"maxAllocationAmount"`
	MaxOpenBuyOrderTime  int      `json:"maxOpenBuyOrderTime"`
	MaxOpenPositions     int      `json:"maxOpenPositions"`
	MaxOpenSellOrderTime int      `json:"maxOpenSellOrderTime"`
	MinOrderAmount       float64  `json:"minOrderAmount"`
	Name                 string   `json:"name"`
	NumberOfTargets      int      `json:"numberOfTargets"`
	ShouldNotifyOnCancel bool     `json:"shouldNotifyOnCancel"`
	ShouldNotifyOnError  bool     `json:"shouldNotifyOnError"`
	ShouldNotifyOnTrade  bool     `json:"shouldNotifyOnTrade"`
	StrategyId           string   `json:"strategy_id"`
	Tickers              []string `json:"tickers"`
}

func (c BotConfig) NewBotFromConfig(strategy Strategy) *Bot {
	return &Bot{
		Config:   c,
		Strategy: strategy,
	}
}

type Bot struct {
	Config   BotConfig
	Strategy Strategy
}

func LoadBotConfig() BotConfig {
	var bot BotConfig

	configFile, err := os.Open("bot-config.json")

	if err != nil {
		log.Fatal(err)
	}

	contentToMarshal, err := ioutil.ReadAll(configFile)

	if err != nil {
		log.Fatal(err)
	}

	json.Unmarshal(contentToMarshal, &bot)

	return bot
}
