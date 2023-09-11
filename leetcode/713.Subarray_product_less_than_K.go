package main

import "fmt"

func numSubarrayProductLessThanK(nums []int, k int) int {
	res := 0
	left, right, mul := 0, 0, 1
	for right < len(nums) {
		mul *= nums[right]
		right++
		if mul < k {
			res += right - left
		} else {
			for mul >= k && left < right {
				mul /= nums[left]
				left++
			}
			res += right - left
		}
	}
	return res
}

func main() {
	nums, k := []int{2,2,2}, 8
	res := numSubarrayProductLessThanK(nums, k)
	fmt.Println(res)
}
