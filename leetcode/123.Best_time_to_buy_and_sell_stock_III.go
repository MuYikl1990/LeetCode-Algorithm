package main

import (
	"fmt"
	"math"
)

func maxProfitIII(prices []int) int {
	firBuy, firSell, secBuy, secSell := math.MinInt, 0, math.MinInt, 0
	for _, price := range prices {
		firBuy = max123(firBuy, -price)
		firSell = max123(firSell, firBuy + price)
		secBuy = max123(secBuy, firSell - price)
		secSell = max123(secSell, secBuy + price)
	}
	return secSell
}

func max123(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	prices := []int{3,3,5,0,0,3,1,4}
	res := maxProfitIII(prices)
	fmt.Println(res)
}
