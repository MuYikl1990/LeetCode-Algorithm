package main

import "fmt"

func lastStoneWeightII(stones []int) int {
	sum, n := 0, len(stones)
	for i := range stones {
		sum += stones[i]
	}
	target := sum / 2
	dp := make([]int, target + 1)
	for i := 0; i < n; i++ {
		for j := target; j >= stones[i]; j-- {
			dp[j] = max1049(dp[j], dp[j - stones[i]] + stones[i])
		}
	}
	return sum - 2 * dp[target]
}

func max1049(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	stones := []int{31,26,33,21,40}
	res := lastStoneWeightII(stones)
	fmt.Println(res)
}
