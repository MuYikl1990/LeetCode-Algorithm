package main

import (
	"fmt"
	"math"
)

func maxProfit(prices []int) int {
	// 成本，利润
	buy, sell := math.MinInt, 0
	for _, price := range prices {
		buy = max121(buy, -price)
		sell = max121(sell, price + buy)
	}
	return sell
}

func max121(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	prices := []int{7,6,2,4,3,1}
	res := maxProfit(prices)
	fmt.Println(res)
}
