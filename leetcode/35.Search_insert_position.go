package main

import "fmt"

func searchInsert(nums []int, target int) int {
	left, right := 0, len(nums)
	for left < right {
		mid := left + (right - left) / 2
		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid
		}
	}

	return left
}

func main() {
	nums, target := []int{1,3,5,6}, 5
	res := searchInsert(nums, target)
	fmt.Println(res)
}
