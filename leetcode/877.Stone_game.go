package main

import "fmt"

func stoneGame(piles []int) bool {
	n := len(piles)
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
		dp[i][i] = piles[i]
	}
	for i := n - 2; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			dp[i][j] = max877(piles[i] - dp[i + 1][j], piles[j] - dp[i][j - 1])
		}
	}
	return dp[0][n - 1] > 0
}

func max877(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	piles := []int{3,7,2,3}
	res := stoneGame(piles)
	fmt.Println(res)
}
