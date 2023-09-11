package main

import "fmt"

func robII(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	var robDp func([]int) int
	robDp = func(nums []int) int {
		dp1, dp2, res := 0, 0, 0
		for i := len(nums) - 1; i >= 0; i-- {
			res = max213(dp1, dp2 + nums[i])
			dp2 = dp1
			dp1 = res
		}
		return res
	}
	return max213(robDp(nums[1:]), robDp(nums[:len(nums) - 1]))
}

func max213(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	nums := []int{2,3,2}
	res := robII(nums)
	fmt.Println(res)
}
