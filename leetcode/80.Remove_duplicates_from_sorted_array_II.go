package main

import "fmt"

func removeDuplicates(nums []int) int {
	length := len(nums)
	if length < 3 {
		return length
	}

	left := 2
	for i := 2; i < length; i++ {
		if nums[i] != nums[left - 2] {
			nums[left] = nums[i]
			left++
		}
	}
	return left
}

func main() {
	nums := []int{0,0,1,1,1,1,2,3,3}
	res := removeDuplicates(nums)
	fmt.Println(res)
}
