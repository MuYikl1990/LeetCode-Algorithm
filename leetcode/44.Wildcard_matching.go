package main

import "fmt"

func isMatchI(s string, p string) bool {
	m, n := len(s), len(p)
	dp := make([][]bool, m + 1)
	for i := range dp {
		dp[i] = make([]bool, n + 1)
	}
	dp[0][0] = true
	for i := 1; i <= n; i++ {
		if p[i - 1] == '*' {
			dp[0][i] = true
		} else {
			break
		}
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if s[i - 1] == p[j - 1] || p[j - 1] == '?' {
				dp[i][j] = dp[i - 1][j - 1]
			} else if p[j - 1] == '*' {
				dp[i][j] = dp[i - 1][j] || dp[i][j - 1]
			}
		}
	}
	return dp[m][n]
}

func main() {
	s, p := "aa", "a"
	res := isMatchI(s, p)
	fmt.Println(res)
}
