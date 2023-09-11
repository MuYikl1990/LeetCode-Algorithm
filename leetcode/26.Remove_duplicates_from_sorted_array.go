package main

import "fmt"

func removeDuplicatesArray(nums []int) int {
	left, right := 0, 0
	for right < len(nums) {
		if nums[left] != nums[right] {
			left++
			nums[left] = nums[right]
		} else {
			right++
		}
	}
	return left + 1
}

func main() {
	nums := []int{0,0,1,1,1,2,2,3,3,4}
	res := removeDuplicatesArray(nums)
	fmt.Println(res)
}
