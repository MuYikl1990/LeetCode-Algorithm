package main

import "fmt"

func findNumberOfLIS(nums []int) int {
	dp, count := make([]int, len(nums)), make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		dp[i] = 1
		count[i] = 1
	}
	max := 1
	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				if dp[j] + 1 > dp[i] {
					dp[i] = dp[j] + 1
					count[i] = count[j]
					max = max673(max, dp[i])
				} else if dp[j] + 1 == dp[i] {
					count[i] += count[j]
				}
			}
		}
	}

	res := 0
	for i := 0; i < len(dp); i++ {
		if dp[i] == max {
			res += count[i]
		}
	}
	return res
}

func max673(a,b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	nums := []int{2,2,2,2,2}
	res := findNumberOfLIS(nums)
	fmt.Println(res)
}
