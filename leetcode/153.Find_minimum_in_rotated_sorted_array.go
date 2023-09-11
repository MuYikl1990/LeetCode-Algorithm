package main

import "fmt"

func findMin(nums []int) int {
	left, right := 0, len(nums) - 1
	for left <= right {
		if nums[left] <= nums[right] {
			return nums[left]
		}
		mid := left + (right - left) / 2
		if nums[left] <= nums[mid] {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return -1
}

func main() {
	nums := []int{4,5,6,7,0,1,2}
	res := findMin(nums)
	fmt.Println(res)
}
