package main

import "fmt"

func superEggDrop(k int, n int) int {
	dp := make([][]int, n + 1)
	for i := range dp {
		dp[i] = make([]int, k + 1)
	}
	for i := 1; i <= n; i++ {
		for j := k; j >= 1; j-- {
			dp[i][j] = dp[i - 1][j - 1] + dp[i - 1][j] + 1
			if dp[i][j] >= n {
				return i
			}
		}
	}
	return n
}

func main() {
	k, n := 2, 100
	res := superEggDrop(k, n)
	fmt.Println(res)
}
