package main

import "fmt"

func longestSubarrayOfOne(nums []int) int {
	res := 0
	left, right, count := 0, 0, 0
	for right < len(nums) {
		if nums[right] == 0 {
			count++
		}
		for count > 1 {
			res = max1493(res, right - left)
			if nums[left] == 0 {
				count--
			}
			left++
		}
		right++
	}
	res = max1493(res, right - left)
	return res - 1
}

func max1493(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	nums := []int{0,1,1,1,0,0,1,1,0,1,1,1,1}
	res := longestSubarrayOfOne(nums)
	fmt.Println(res)
}
