package main

import (
	"fmt"
	"sort"
)

func triangleNumber(nums []int) int {
	sort.Ints(nums)
	res := 0

	for i := len(nums) - 1; i > 1; i-- {
		left, right := 0, i - 1
		for left < right {
			if nums[left] + nums[right] > nums[i] {
				res += right - left
				right--
			} else {
				left++
			}
		}
	}
	return res
}

func main() {
	nums := []int{4,2,3,4}
	res := triangleNumber(nums)
	fmt.Println(res)
}
