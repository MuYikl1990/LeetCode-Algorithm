package main

import (
	"fmt"
	"math"
)

func maxProfitIV(k int, prices []int) int {
	buy, sell := make([]int, k + 1), make([]int, k + 1)
	for i := 0; i <= k; i++ {
		buy[i] = math.MinInt
		sell[i] = 0
	}

	for _, price := range prices {
		for i := 1; i <= k; i++ {
			buy[i] = max188(buy[i], sell[i - 1] - price)
			sell[i] = max188(sell[i], buy[i] + price)
		}
	}
	return sell[k]
}

func max188(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	k, prices := 2, []int{3,2,6,5,0,3}
	res := maxProfitIV(k, prices)
	fmt.Println(res)
}
