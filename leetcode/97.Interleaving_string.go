package main

import "fmt"

func isInterleave(s1 string, s2 string, s3 string) bool {
	m, n, l := len(s1), len(s2), len(s3)
	if m + n != l {
		return false
	}
	dp := make([][]bool, m + 1)
	for i := range dp {
		dp[i] = make([]bool, n + 1)
	}
	dp[0][0] = true
	for i := 1; i <= m && s1[i - 1] == s3[i - 1]; i++ {
		dp[i][0] = true
	}
	for j := 1; j <= n && s2[j - 1] == s3[j - 1]; j++ {
		dp[0][j] = true
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			dp[i][j] = (dp[i - 1][j] && s1[i - 1] == s3[i + j - 1]) || (dp[i][j - 1] && s2[j - 1] == s3[i + j - 1])
		}
	}
	return dp[m][n]
}

func main() {
	s1, s2, s3 := "aabcc", "dbbca", "aadbbabcac"
	res := isInterleave(s1, s2, s3)
	fmt.Println(res)
}
