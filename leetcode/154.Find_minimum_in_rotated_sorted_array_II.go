package main

import "fmt"

func findMinII(nums []int) int {
	left, right := 0, len(nums) - 1
	for left < right {
		mid := left + (right - left) / 2
		if nums[mid] > nums[right] {
			left = mid + 1
		} else if nums[mid] < nums[right] {
			right = mid
		} else {
			right--
		}
	}
	return nums[left]
}

func main() {
	nums := []int{2,2,2,0,1}
	res := findMinII(nums)
	fmt.Println(res)
}
