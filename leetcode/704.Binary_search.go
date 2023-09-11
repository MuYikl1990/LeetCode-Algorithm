package main

import "fmt"

func searchI(nums []int, target int) int {
	left, right := 0, len(nums) - 1
	for left <= right {
		mid := left + (right - left) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}

func main() {
	nums, target := []int{-1,0,3,5,9,12}, 9
	res := searchI(nums, target)
	fmt.Println(res)
}
