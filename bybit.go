package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// BybitTickerData represents the structure of Bybit's ticker response
type BybitTickerData struct {
	RetCode int `json:"retCode"`
	Result  struct {
		List []struct {
			Symbol         string `json:"symbol"`
			LastPrice      string `json:"lastPrice"`
			High24h        string `json:"highPrice24H"`
			Low24h         string `json:"lowPrice24H"`
			Volume24h      string `json:"volume24H"`
			PriceChange24h string `json:"price24HPcnt"`
		} `json:"list"`
	} `json:"result"`
}

func getBybitBTCData() (*BybitTickerData, error) {
	resp, err := http.Get("https://api.bybit.com/v5/market/tickers?category=spot&symbol=BTCUSDT")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var ticker BybitTickerData
	if err := json.Unmarshal(body, &ticker); err != nil {
		return nil, fmt.Errorf("error unmarshaling Bybit response: %v", err)
	}

	return &ticker, nil
}
