package main

import (
	"fmt"
)

func minOperations(nums []int, x int) int {
	n := len(nums)
	if nums[0] > x && nums[n - 1] > x {
		return -1
	}

	sum := 0
	for _, num := range nums {
		sum += num
	}
	target, res := sum - x, -1
	if target < 0 {
		return -1
	}
	left, right := 0, 0
	sum = 0
	for right < n {
		sum += nums[right]
		for sum > target {
			sum -= nums[left]
			left++
		}
		if sum == target {
			res = max1658(res, right - left + 1)
		}
		right++
	}
	if res == -1 {
		return -1
	}
	return n - res
}

func max1658(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	nums, x := []int{1,1}, 3
	res := minOperations(nums, x)
	fmt.Println(res)
}
