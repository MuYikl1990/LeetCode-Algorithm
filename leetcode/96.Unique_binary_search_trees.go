package main

import "fmt"

func numTrees(n int) int {
	dp := make([]int, n + 1)
	dp[0], dp[1] = 1, 1

	for i := 2; i < n + 1; i++ {
		for j := 0; i - j > 0; j++ {
			dp[i] += dp[j] * dp[i - j - 1]
		}
	}
	return dp[n]
}

func main() {
	n := 3
	res := numTrees(n)
	fmt.Println(res)
}
