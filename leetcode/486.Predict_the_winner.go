package main

import "fmt"

func PredictTheWinner(nums []int) bool {
	length := len(nums)
	dp := make([][]int, length)
	for i := range dp {
		dp[i] = make([]int, length)
	}

	for i := range nums {
		dp[i][i] = nums[i]
	}

	for i := length - 2; i >= 0; i-- {
		for j := i + 1; j < length; j++ {
			dp[i][j] = max486(nums[i] - dp[i + 1][j], nums[j] - dp[i][j - 1])
		}
	}
	return dp[0][length - 1] >= 0
}

func max486(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	nums := []int{1,5,233,7}
	res := PredictTheWinner(nums)
	fmt.Println(res)
}
