package main

import "fmt"

func maxSum(nums []int) int {
	nums = append(nums, 0)
	nums = append([]int{0}, nums...)
	preSum := make([]int, len(nums) + 1)
	for i := range nums {
		preSum[i + 1] = preSum[i] + nums[i]
	}

	var stack []int
	res := 0
	for i := 0; i < len(nums); i++ {
		for len(stack) > 0 && nums[i] < nums[stack[len(stack) - 1]] {
			pos := stack[len(stack) - 1]
			stack = stack[:len(stack) - 1]
			sum := preSum[i] - preSum[stack[len(stack) - 1] + 1]
			res = max03(res, sum * nums[pos])
		}
		stack = append(stack, i)
	}
	return res
}

func max03(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	nums := []int{6,2,1}
	res := maxSum(nums)
	fmt.Println(res)
}
