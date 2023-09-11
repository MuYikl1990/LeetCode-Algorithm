package main

import "fmt"

func nextPermutation(nums []int)  {
	if len(nums) < 2 {
		return
	}

	left, right, end := len(nums) - 2, len(nums) - 1, len(nums) - 1

	for left >= 0 && nums[left] >= nums[right] {
		left--
		right--
	}

	if left >= 0 {
		for nums[left] >= nums[end] {
			end--
		}
		nums[left], nums[end] = nums[end], nums[left]
	}

	for i, j := right, len(nums) - 1; i < j; i, j = i + 1, j - 1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
}

func main() {
	num := []int{1,1,5}
	nextPermutation(num)
	fmt.Println(num)
}
