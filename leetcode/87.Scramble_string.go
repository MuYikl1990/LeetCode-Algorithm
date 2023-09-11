package main

import "fmt"

func isScramble(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	n := len(s1)
	dp := make([][][]bool, n)
	for i := range dp {
		dp[i] = make([][]bool, n)
		for j := range dp[i] {
			dp[i][j] = make([]bool, n + 1)
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			dp[i][j][1] = s1[i] == s2[j]
		}
	}

	for length := 2; length <= n; length++ {
		for i := 0; i <= n - length; i++ {
			for j := 0; j <= n - length; j ++ {
				for k := 1; k <= length - 1; k++ {
					if dp[i][j][k] && dp[i + k][j + k][length - k] {
						dp[i][j][length] = true
						break
					}
					if dp[i][j + length - k][k] && dp[i + k][j][length - k] {
						dp[i][j][length] = true
						break
					}
				}
			}
		}
	}
	return dp[0][0][n]
}

func main() {
	s1, s2 := "great", "rgeat"
	res := isScramble(s1, s2)
	fmt.Println(res)
}
