package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

type SignalComparisonAndValue struct {
	Comparison string `json:"comparison"`
	Value      int    `json:"value"`
}

type StrategySegmentParams struct {
	Period     int                      `json:"period"`
	SignalWhen SignalComparisonAndValue `json:"signal_when"`
}

type StrategySegment struct {
	Indicator        string                `json:"indicator"`
	IsRequired       bool                  `json:"is_required"`
	KeepSignalLength int                   `json:"keep_signal_length"`
	Params           StrategySegmentParams `json:"params"`
	Period           int                   `json:"period"`
	StrategyType     string                `json:"strategy_type"`
}

type Strategy struct {
	Exchange   string            `json:"exchange"`
	Id         string            `json:"id"`
	Name       string            `json:"name"`
	Strategies []StrategySegment `json:"strategies"`
}

func LoadStrategy() Strategy {
	var strategy Strategy

	configFile, err := os.Open("strategy-config.json")

	if err != nil {
		log.Fatal(err)
	}

	contentToMarshal, err := ioutil.ReadAll(configFile)

	if err != nil {
		log.Fatal(err)
	}

	json.Unmarshal(contentToMarshal, &strategy)

	return strategy
}
