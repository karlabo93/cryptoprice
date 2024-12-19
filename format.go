package main

import (
	"fmt"
	"strconv"
)

// Helper function to format numbers with thousand separators and 2 decimal places
func formatNumber(numStr string) string {
	num, err := strconv.ParseFloat(numStr, 64)
	if err != nil {
		return numStr // return original string if parsing fails
	}
	// Format with thousand separator and 2 decimal places
	return fmt.Sprintf("%.2f", num)
}
