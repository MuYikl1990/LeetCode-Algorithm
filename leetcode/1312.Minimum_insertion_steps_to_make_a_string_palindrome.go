package main

import "fmt"

func minInsertions(s string) int {
	dp := make([][]int, len(s))
	for i := range dp {
		dp[i] = make([]int, len(s))
		dp[i][i] = 1
	}
	for i := len(s) - 1; i >= 0; i-- {
		for j := i + 1; j < len(s); j++ {
			if s[i] == s[j] {
				dp[i][j] = dp[i + 1][j - 1] + 2
			} else {
				dp[i][j] = max1312(dp[i + 1][j], dp[i][j - 1])
			}
		}
	}
	return len(s) - dp[0][len(s) - 1]
}

func max1312(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	s := "leetcode"
	res := minInsertions(s)
	fmt.Println(res)
}
