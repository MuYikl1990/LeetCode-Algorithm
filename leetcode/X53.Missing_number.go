package main

import "fmt"

func missingNumber(nums []int) int {
	left, right := 0, len(nums) - 1

	for left <= right {
		mid := left + (right - left) / 2
		if mid >= nums[mid] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return left
}

func main() {
	nums := []int{0}
	res := missingNumber(nums)
	fmt.Println(res)
}
