package main

import (
	"fmt"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	res := nums[0] + nums[1] + nums[2]
	for i := 0; i < len(nums) - 1; i++ {
		left, right := i + 1, len(nums) - 1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if abs16(sum, target) < abs16(res, target) {
				res = sum
			}
			if sum > target {
				right--
			} else if sum < target {
				left++
			} else {
				return res
			}
		}
	}
	return res
}

func abs16(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

func main() {
	nums, target := []int{-1,2,1,-4}, 1
	res := threeSumClosest(nums, target)
	fmt.Println(res)
}
