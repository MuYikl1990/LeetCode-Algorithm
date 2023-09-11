package main

import "fmt"

func minDistance(word1 string, word2 string) int {
	m, n := len(word1), len(word2)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	for i := 0; i <= m; i++ {
		dp[i][0] = i
	}
	for j := 0; j <= n; j++ {
		dp[0][j] = j
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min72(dp[i-1][j-1] + 1, min72(dp[i][j-1] + 1, dp[i-1][j] + 1))
			}
		}
	}
	return dp[m][n]
}

func min72(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func main() {
	word1, word2 := "intention", "execution"
	res := minDistance(word1, word2)
	fmt.Println(res)
}
