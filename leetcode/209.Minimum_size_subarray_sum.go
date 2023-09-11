package main

import (
	"fmt"
	"math"
)

func minSubArrayLen(target int, nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	left, right, res, sum := 0, 0, math.MaxInt, 0
	for right < n {
		sum += nums[right]
		right++
		for sum >= target {
			res = min209(res, right - left)
			sum -= nums[left]
			left++
		}
	}
	if res == math.MaxInt {
		return 0
	}
	return res
}

func min209(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	target, nums := 7, []int{2,3,-1,2,4,-3}
	res := minSubArrayLen(target, nums)
	fmt.Println(res)
}
