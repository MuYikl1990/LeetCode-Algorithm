package main

import "fmt"

func longestCommonSubsequence(text1 string, text2 string) int {
	t1Len := len(text1)
	t2Len := len(text2)
	dp := make([][]int, t1Len + 1)
	for i := range dp {
		dp[i] = make([]int, t2Len + 1)
	}

	for i := 1; i <= t1Len; i++ {
		for j := 1; j <= t2Len; j++ {
			if text1[i - 1] == text2[j - 1] {
				dp[i][j] = dp[i - 1][j - 1] + 1
			} else {
				dp[i][j] = max1143(dp[i][j - 1], dp[i - 1][j])
			}
		}
	}
	return dp[t1Len][t2Len]
}

func max1143(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	text1, text2 := "abcde", "ace"
	res := longestCommonSubsequence(text1, text2)
	fmt.Println(res)
}
