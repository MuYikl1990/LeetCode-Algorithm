package main

import (
	"fmt"
	"math"
)

func maxDotProduct(nums1 []int, nums2 []int) int {
	m, n := len(nums1), len(nums2)
	dp := make([][]int, m + 1)
	for i := range dp {
		dp[i] = make([]int, n + 1)
		for j := range dp[i] {
			dp[i][j] = math.MinInt8
		}
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			dp[i][j] = nums1[i - 1] * nums2[j - 1]
			dp[i][j] = max1458(dp[i][j], nums1[i - 1] * nums2[j - 1] + dp[i - 1][j - 1])
			dp[i][j] = max1458(dp[i][j], dp[i - 1][j])
			dp[i][j] = max1458(dp[i][j], dp[i][j - 1])
		}
	}
	return dp[m][n]
}

func max1458(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	nums1, nums2 := []int{-1,-1}, []int{1,1}
	res := maxDotProduct(nums1, nums2)
	fmt.Println(res)
}
