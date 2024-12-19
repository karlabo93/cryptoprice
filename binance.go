package main

import (
	"encoding/json"
	"io"
	"net/http"
)

// BinanceTickerData represents the structure of Binance's ticker response
type BinanceTickerData struct {
	Symbol             string `json:"symbol"`
	PriceChange        string `json:"priceChange"`
	PriceChangePercent string `json:"priceChangePercent"`
	LastPrice          string `json:"lastPrice"`
	Volume             string `json:"volume"`
	High24h            string `json:"highPrice"`
	Low24h             string `json:"lowPrice"`
}

func getBinanceBTCData() (*BinanceTickerData, error) {
	resp, err := http.Get("https://api.binance.com/api/v3/ticker/24hr?symbol=BTCUSDT")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var ticker BinanceTickerData
	if err := json.Unmarshal(body, &ticker); err != nil {
		return nil, err
	}

	return &ticker, nil
}
