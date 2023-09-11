package main

import "fmt"

func numRollsToTarget(n int, k int, target int) int {
	dp := make([][]int, n + 1)
	for i := range dp {
		dp[i] = make([]int, target + 1)
	}
	dp[0][0] = 1

	for i := 1; i <= n; i++ {
		for j := 1; j <= target; j++ {
			for m := 1; m <= k; m++ {
				if j >= m {
					dp[i][j] = (dp[i][j] + dp[i - 1][j - m]) % 1000000007
				}
			}
		}
	}
	return dp[n][target]
}

func numRollsToTarget1(n int, k int, target int) int {
	dp := make([]int, target + 1)
	dp[0] = 1

	for i := 1; i <= n; i++ {
		for j := target; j >= 0; j-- {
			dp[j] = 0
			for m := 1; m <= k; m++ {
				if j >= m {
					dp[j] = (dp[j] + dp[j - m]) % 1000000007
				}
			}
		}
	}
	return dp[target]
}

func main() {
	n, k, target := 30, 30, 500
	res := numRollsToTarget1(n, k, target)
	fmt.Println(res)
}
