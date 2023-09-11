package main

import "fmt"

func longestOnes(nums []int, k int) int {
	n := len(nums)
	left, right := 0, 0
	res := 0
	for right < n {
		if nums[right] == 0 {
			k--
		}
		for k < 0 {
			if nums[left] == 0 {
				k++
			}
			left++
		}
		res = max1004(res, right - left + 1)
		right++
	}
	return res
}

func max1004(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	nums, k := []int{0,0,1,1,0,0,1,1,1,0,1,1,0,0,0,1,1,1,1}, 3
	res := longestOnes(nums, k)
	fmt.Println(res)
}
