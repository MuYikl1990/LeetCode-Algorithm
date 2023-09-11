package main

import "fmt"

func longestPalindromeSubseq(s string) int {
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
				dp[i][j] = max516(dp[i][j - 1], dp[i + 1][j])
			}
		}
	}
	return dp[0][len(s) - 1]
}

func max516(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	s := "leetcode"
	res := longestPalindromeSubseq(s)
	fmt.Println(res)
}
