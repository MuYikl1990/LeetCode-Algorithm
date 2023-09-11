package main

import "fmt"

func findPeakElement(nums []int) int {
	left, right := 0, len(nums) - 1
	for left < right {
		mid := left + (right - left) / 2
		if nums[mid] > nums[mid + 1] {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return right
}

func main() {
	nums := []int{1,2,1,3,5,6,4}
	res := findPeakElement(nums)
	fmt.Println(res)
}
