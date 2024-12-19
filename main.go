package main

import (
	"fmt"
	"strconv"
	"time"
)

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
		vol, _ := strconv.ParseFloat(data.PriceChange24h, 64)
		vol *= 100
		fmt.Printf("24h Change: %.2f%%\n", vol)
		fmt.Printf("24h High: $%s\n", formatNumber(data.High24h))
		fmt.Printf("24h Low: $%s\n", formatNumber(data.Low24h))
		fmt.Printf("24h Volume: %s BTC\n", formatNumber(data.Volume24h))
	} else {
		fmt.Println("\nNo Bybit data available")
	}
}
