package main

import "fmt"

func numDistinct(s string, t string) int {
	sLen := len(s)
	tLen := len(t)
	if sLen < tLen {
		return 0
	}
	dp := make([][]int, sLen + 1)
	for i := range dp {
		dp[i] = make([]int, tLen + 1)

	}

	for i := 0; i <= tLen; i++ {
		dp[0][i] = 0
	}

	for i := 0; i <= sLen; i++ {
		dp[i][0] = 1
	}

	for i := 1; i <= sLen; i++ {
		for j := 1; j <= tLen; j++ {
			if s[i - 1] == t[j - 1] {
				dp[i][j] = dp[i-1][j-1] + dp[i-1][j]
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}
	return dp[sLen][tLen]
}

func main() {
	s, t := "rabbbit", "rabbit"
	res := numDistinct(s, t)
	fmt.Println(res)
}
