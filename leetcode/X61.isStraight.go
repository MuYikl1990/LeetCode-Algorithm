package main

import (
	"fmt"
)

func isStraight(nums []int) bool {
	min, max := 15, -1
	arr := [14]int{}
	for i := range nums {
		if nums[i] == 0 {
			arr[0]++
			continue
		}
		if arr[nums[i]] != 0 {
			return false
		}
		arr[nums[i]]++
		if nums[i] < min {
			min = nums[i]
		}
		if nums[i] > max {
			max = nums[i]
		}
		if max - min >= 5 {
			return false
		}
	}
	return true
}

func main() {
	nums := []int{1,6,5,4,2}
	res := isStraight(nums)
	fmt.Println(res)
}
