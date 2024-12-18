package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
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

// Helper function to format numbers with thousand separators and 2 decimal places
func formatNumber(numStr string) string {
	num, err := strconv.ParseFloat(numStr, 64)
	if err != nil {
		return numStr // return original string if parsing fails
	}
	// Format with thousand separator and 2 decimal places
	return fmt.Sprintf("%.2f", num)
}

func main() {
	// Fetch Binance data
	binanceTicker, err := getBinanceBTCData()
	if err != nil {
		fmt.Printf("Error fetching Binance data: %v\n", err)
	} else {
		fmt.Println("\nBinance Bitcoin Information:")
		fmt.Printf("Current Price: $%s\n", formatNumber(binanceTicker.LastPrice))
		fmt.Printf("24h Change: %s%%\n", formatNumber(binanceTicker.PriceChangePercent))
		fmt.Printf("24h High: $%s\n", formatNumber(binanceTicker.High24h))
		fmt.Printf("24h Low: $%s\n", formatNumber(binanceTicker.Low24h))
		fmt.Printf("24h Volume: %s BTC\n", formatNumber(binanceTicker.Volume))
	}

	// Add a small delay between requests
	time.Sleep(100 * time.Millisecond)

	// Fetch Bybit data
	bybitTicker, err := getBybitBTCData()
	if err != nil {
		fmt.Printf("Error fetching Bybit data: %v\n", err)
	} else if len(bybitTicker.Result.List) > 0 {
		data := bybitTicker.Result.List[0]
		fmt.Println("\nBybit Bitcoin Information:")
		fmt.Printf("Current Price: $%s\n", formatNumber(data.LastPrice))
		fmt.Printf("24h Change: %s%%\n", formatNumber(data.PriceChange24h))
		fmt.Printf("24h High: $%s\n", formatNumber(data.High24h))
		fmt.Printf("24h Low: $%s\n", formatNumber(data.Low24h))
		fmt.Printf("24h Volume: %s BTC\n", formatNumber(data.Volume24h))
	} else {
		fmt.Println("\nNo Bybit data available")
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
