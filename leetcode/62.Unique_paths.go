package main

import "fmt"

func uniquePaths(m int, n int) int {
	dp := make([]int, n)
	for k := 0; k < n; k++ {
		dp[k] = 1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[j] = dp[j] + dp[j - 1]
		}
	}
	return dp[n-1]
}

func main() {
	m, n := 7, 3
	res := uniquePaths(m , n)
	fmt.Println(res)
}
