package main

import "fmt"

func coinChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}

	dp := make([]int, amount + 1)
	for i := range dp {
		if i == 0 {
			dp[i] = 0
		} else {
			dp[i] = amount + 1
		}
	}

	for j := 0; j < len(dp); j++ {
		for _, coin := range coins {
			if j - coin < 0 {
				continue
			}
			dp[j] = min(dp[j], 1 + dp[j - coin])
		}
	}
	if dp[amount] == amount + 1 {
		return -1
	}
	return dp[amount]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	coins, amount := []int{1,2,5}, 11
	res := coinChange(coins, amount)
	fmt.Println(res)
}
