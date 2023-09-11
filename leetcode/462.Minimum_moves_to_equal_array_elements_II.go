package main

import (
	"fmt"
	"sort"
)

func minMoves2(nums []int) int {
	sort.Ints(nums)
	target, res := 0, 0
	if len(nums) % 2 == 0 {
		target = (nums[len(nums) / 2 - 1] + nums[len(nums) / 2]) / 2
	} else {
		target = nums[len(nums) / 2]
	}

	for j := range nums {
		if target > nums[j] {
			res += target - nums[j]
		} else {
			res += nums[j] - target
		}
	}
	return res
}

func main() {
	nums := []int{1}
	res := minMoves2(nums)
	fmt.Println(res)
}
