package main

import "fmt"

func sortedSquares(nums []int) []int {
	res := make([]int, len(nums))
	left, right := 0, len(nums) - 1
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[left] * nums[left] < nums[right] * nums[right] {
			res[i] = nums[right] * nums[right]
			right--
		} else {
			res[i] = nums[left] * nums[left]
			left++
		}
	}
	return res
}

func main() {
	nums := []int{-7,-3,2,3,11}
	res := sortedSquares(nums)
	fmt.Println(res)
}
