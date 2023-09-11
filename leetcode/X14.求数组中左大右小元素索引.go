package main

import (
	"fmt"
	"math"
)

// 求数组中比左边元素都大同时比右边元素都小的元素，返回这些元素的索引
func findIndex(nums []int) []int {
	leftMax, rightMin := make([]int, len(nums)), make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		leftMax[i] = math.MinInt
		rightMin[i] = math.MaxInt
	}

	for i := 1; i < len(nums); i++ {
		leftMax[i] = max14(leftMax[i - 1], nums[i - 1])
	}
	for i := len(nums) - 2; i >= 0; i-- {
		rightMin[i] = min14(rightMin[i + 1], nums[i + 1])
	}

	var res []int
	for i := 0; i < len(nums); i++ {
		if leftMax[i] < nums[i] && nums[i] < rightMin[i] {
			res = append(res, i)
		}
	}
	return res
}

func max14(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min14(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	nums := []int{2, 3, 1, 8, 9, 20, 12}
	res := findIndex(nums)
	fmt.Println(res)
}
