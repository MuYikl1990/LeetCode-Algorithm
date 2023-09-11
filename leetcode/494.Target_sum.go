package main

import "fmt"

func findTargetSumWays(nums []int, target int) int {
	sum := 0
	for i := range nums {
		sum += nums[i]
	}

	if sum < target || (sum - target) % 2 != 0 {
		return 0
	}

	// 一维数组
	neg := (sum - target) / 2
	dp := make([]int, neg + 1)
	dp[0] = 1

	for i := 0; i < len(nums); i++ {
		num := nums[i]
		for j := neg; j >= num; j-- {
			dp[j] += dp[j - num]
		}
	}
	return dp[neg]

	// 二维数组
	//dp, neg := make([][]int, len(nums) + 1), (sum - target) / 2
	//for i := range dp {
	//	dp[i] = make([]int, neg + 1)
	//}
	//
	//dp[0][0] = 1
	//for i := 1; i <= len(nums); i++ {
	//	num := nums[i - 1]
	//	for j := 0; j <= neg; j++ {
	//		dp[i][j] = dp[i - 1][j]
	//		if num <= j {
	//			dp[i][j] += dp[i - 1][j - num]
	//		}
	//	}
	//}
	//return dp[len(nums)][neg]
}

func main() {
	nums, target := []int{1,1,1,1,1}, 3
	res := findTargetSumWays(nums, target)
	fmt.Println(res)
}
