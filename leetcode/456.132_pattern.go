package main

import (
	"fmt"
	"math"
)

func find132pattern(nums []int) bool {
	nums = append([]int{math.MaxInt}, nums...)
	minList := make([]int, len(nums))
	minNum := nums[0]
	for i := 1; i < len(nums); i++ {
		minNum = min456(minNum, nums[i - 1])
		minList[i] = minNum
	}

	var stack []int
	for i := len(nums) - 1; i > 0; i-- {
		tmp := math.MinInt
		for len(stack) > 0 && nums[i] > nums[stack[len(stack) - 1]] {
			tmp = nums[stack[len(stack) - 1]]
			stack = stack[:len(stack) - 1]
		}
		if minList[i] < tmp {
			return true
		}
		stack = append(stack, i)
	}
	return false
}

func min456(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	nums := []int{1,0,1,-4,-3}
	res := find132pattern(nums)
	fmt.Println(res)
}
