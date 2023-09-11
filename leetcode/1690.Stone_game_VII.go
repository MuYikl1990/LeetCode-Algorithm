package main

import "fmt"

func stoneGameVII(stones []int) int {
	preSum := make([]int, len(stones) + 1)
	dp := make([][]int, len(stones))
	for i := range dp {
		dp[i] = make([]int, len(stones))
	}

	for i := range stones {
		preSum[i + 1] = preSum[i] + stones[i]
	}

	for i := len(stones) - 1; i >= 0; i-- {
		for j := i + 1; j < len(stones); j++ {
			dp[i][j] = max1690(preSum[j + 1] - preSum[i + 1] - dp[i + 1][j], preSum[j] - preSum[i] - dp[i][j - 1])
		}
	}
	return dp[0][len(stones) - 1]
}

func max1690(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	stones := []int{7,90,5,1,100,10,10,2}
	res := stoneGameVII(stones)
	fmt.Println(res)
}
