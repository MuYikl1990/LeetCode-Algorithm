package main

import "fmt"

func isMatch(s string, p string) bool {
	m, n := len(s), len(p)
	dp := make([][]bool, m + 1)
	for i := range dp {
		dp[i] = make([]bool, n + 1)
	}
	dp[0][0] = true
	for i := 1; i < n; i++ {
		if p[i - 1] == '*' {
			dp[0][i] = dp[0][i-2]
		}
	}

	for i := 1; i < m + 1; i++ {
		for j := 1; j < n + 1; j++ {
			if s[i-1] == p[j-1] || p[j-1] == '.' {
				dp[i][j] = dp[i-1][j-1]
			} else if p[j-1] == '*' {
				if s[i-1] == p[j-2] || p[j-2] == '.' {
					dp[i][j] = dp[i-1][j-2] || dp[i][j-2] || dp[i-1][j]
				} else {
					dp[i][j] = dp[i][j-2]
				}
			}
		}
	}
	return dp[m][n]
}

func main() {
	s, p := "aba", "a*"
	res := isMatch(s, p)
	fmt.Println(res)
}
