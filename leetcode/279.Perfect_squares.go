package main

import (
	"fmt"
	"math"
)

func numSquares(n int) int {
	dp := make([]int, n + 1)
	for i := range dp {
		dp[i] = math.MaxInt
	}
	dp[0] = 0
	for i := 1; i <= n; i++ {
		for j := 1; j * j <= i; j++ {
			dp[i] = min279(dp[i], dp[i - j * j] + 1)
		}
	}
	return dp[n]
}

func min279(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	n := 12
	res := numSquares(n)
	fmt.Println(res)
}
