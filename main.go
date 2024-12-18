package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
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

func main() {
	// Fetch Binance data
	binanceTicker, err := getBinanceBTCData()
	if err != nil {
		fmt.Printf("Error fetching Binance data: %v\n", err)
	} else {
		fmt.Printf("\nBinance Bitcoin Information:\n")
		fmt.Printf("Current Price: $%s\n", binanceTicker.LastPrice)
		fmt.Printf("24h Change: %s%%\n", binanceTicker.PriceChangePercent)
		fmt.Printf("24h High: $%s\n", binanceTicker.High24h)
		fmt.Printf("24h Low: $%s\n", binanceTicker.Low24h)
		fmt.Printf("24h Volume: %s BTC\n", binanceTicker.Volume)
	}

	// Add a small delay between requests
	time.Sleep(100 * time.Millisecond)

	// Fetch Bybit data
	bybitTicker, err := getBybitBTCData()
	if err != nil {
		fmt.Printf("Error fetching Bybit data: %v\n", err)
	} else if len(bybitTicker.Result.List) > 0 {
		data := bybitTicker.Result.List[0]
		fmt.Printf("\nBybit Bitcoin Information:\n")
		fmt.Printf("Current Price: $%s\n", data.LastPrice)
		fmt.Printf("24h Change: %s%%\n", data.PriceChange24h)
		fmt.Printf("24h High: $%s\n", data.High24h)
		fmt.Printf("24h Low: $%s\n", data.Low24h)
		fmt.Printf("24h Volume: %s BTC\n", data.Volume24h)
	} else {
		fmt.Printf("\nNo Bybit data available\n")
	}
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
