package main

import "fmt"

func maxWidthRamp(nums []int) int {
	res := 0
	var stack []int
	for i := 0; i < len(nums); i++ {
		if len(stack) == 0 || nums[stack[len(stack) - 1]] > nums[i] {
			stack = append(stack, i)
		}
	}

	for j := len(nums) - 1; j >= 0; j-- {
		for len(stack) != 0 && nums[stack[len(stack) - 1]] <= nums[j] {
			pos := stack[len(stack) - 1]
			stack = stack[:len(stack) - 1]
			res = max962(res, j - pos)
		}
	}
	return res
}

func max962(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	nums := []int{9,8,1,0,1,9,4,0,4,1}
	res := maxWidthRamp(nums)
	fmt.Println(res)
}
