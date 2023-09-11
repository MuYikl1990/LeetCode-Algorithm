package main

import "fmt"

func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	left, right := -1, -1
	start, end := 0, len(nums) - 1
	for start < end {
		mid := start + (end - start) / 2
		if nums[mid] >= target {
			end = mid
		} else {
			start = mid + 1
		}
	}
	if nums[end] != target {
		return []int{-1, -1}
	}
	left, start, end = end, end, len(nums) - 1
	for start < end {
		mid := start + (end - start + 1) / 2
		if nums[mid] <= target {
			start = mid
		} else {
			end = mid - 1
		}
	}
	right = end
	return []int{left, right}
}

func main() {
	nums, target := []int{5,7,7,8,8,10}, 7
	res := searchRange(nums, target)
	fmt.Println(res)
}
