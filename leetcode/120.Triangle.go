package main

import (
	"fmt"
)

func minimumTotal(triangle [][]int) int {
	dp := make([]int, len(triangle) + 1)
	for i := len(triangle) - 1; i >= 0; i-- {
		for j := 0; j < len(triangle[i]); j++ {
			dp[j] = min120(dp[j], dp[j + 1]) + triangle[i][j]
		}
	}
	return dp[0]
}

func min120(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	triangle := [][]int{{-1}, {2,3}, {1,-1,-3}}
	res := minimumTotal(triangle)
	fmt.Println(res)
}
