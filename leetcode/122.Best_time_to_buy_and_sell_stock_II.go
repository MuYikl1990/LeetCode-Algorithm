package main

import (
	"fmt"
	"math"
)

func maxProfitII(prices []int) int {
	buy, sell := math.MinInt, 0
	for _, price := range prices {
		buy = max122(buy, sell - price)
		sell = max122(sell, price + buy)
	}
	return sell
}

func max122(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	prices := []int{7,1,5,3,6,9}
	res := maxProfitII(prices)
	fmt.Println(res)
}
