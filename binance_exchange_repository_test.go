package main

import "testing"

func TestBinanceExchangeRepository_IsSymbolTradable(t *testing.T) {
	t.Run("When passed a tradable symbol and empty ExchangeInfoSymbol list", func(t *testing.T) {
		e := BinanceExchangeRepository{}
		exchange_info_mock := []ExchangeInfoSymbol{}
		found := e.IsSymbolTradable("BTCUSDT", exchange_info_mock)
		expectedResult := false

		if found != expectedResult {
			t.Errorf("BinanceExchangeRepository.IsSymbolTradable() = %v, want %v", found, expectedResult)
		}
	})

	t.Run("When passed a tradable symbol that is not present in ExchangeInfoSymbol", func(t *testing.T) {
		e := BinanceExchangeRepository{}
		exchange_info_mock := []ExchangeInfoSymbol{{"BTCUSDT", "TRADING", "BTC", "USDT"}}
		found := e.IsSymbolTradable("AMD", exchange_info_mock)
		expectedResult := false

		if found != expectedResult {
			t.Errorf("BinanceExchangeRepository.IsSymbolTradable() = %v, want %v", found, expectedResult)
		}
	})

	t.Run("When passed a tradable symbol that is present in ExchangeInfoSymbol", func(t *testing.T) {
		t.Run("and status is TRADING", func(t *testing.T) {
			e := BinanceExchangeRepository{}
			exchange_info_mock := []ExchangeInfoSymbol{{"BTCUSDT", "TRADING", "BTC", "USDT"}}
			found := e.IsSymbolTradable("BTCUSDT", exchange_info_mock)
			expectedResult := true

			if found != expectedResult {
				t.Errorf("BinanceExchangeRepository.IsSymbolTradable() = %v, want %v", found, expectedResult)
			}
		})

		t.Run("and status is not TRADING", func(t *testing.T) {
			e := BinanceExchangeRepository{}
			exchange_info_mock := []ExchangeInfoSymbol{{"BTCUSDT", "threeve", "BTC", "USDT"}}
			found := e.IsSymbolTradable("BTCUSDT", exchange_info_mock)
			expectedResult := false

			if found != expectedResult {
				t.Errorf("BinanceExchangeRepository.IsSymbolTradable() = %v, want %v", found, expectedResult)
			}
		})
	})
}
