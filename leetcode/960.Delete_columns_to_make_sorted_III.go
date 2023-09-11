package main

import "fmt"

func minDeletionSizeIII(strs []string) int {
	m, n := len(strs), len(strs[0])
	dp := make([]int, n)
	for i := range dp {
		dp[i] = 1
	}

	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			flag := true
			for k := 0; k < m; k++ {
				if strs[k][j] > strs[k][i] {
					flag = false
					break
				}
			}
			if flag {
				dp[i] = max960(dp[i], dp[j] + 1)
			}
		}
	}

	res := 0
	for i := range dp {
		res = max960(res, dp[i])
	}
	return n - res
}

func max960(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	strs := []string{"babca","bbazb"}
	res := minDeletionSizeIII(strs)
	fmt.Println(res)
}
