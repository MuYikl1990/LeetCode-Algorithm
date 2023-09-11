package main

import "fmt"

func missingElement(nums []int, k int) int {
	left, right := 0, len(nums) - 1
	if nums[right] - nums[left] - right < k {
		return nums[right] + k - (nums[right] - nums[left] - right)
	}

	for left < right {
		mid := left + (right - left) / 2
		diff := nums[mid] - nums[0] - mid
		if diff >= k {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return nums[left - 1] + k - (nums[left - 1] - nums[0] - left + 1)
}

func main() {
	nums, k := []int{4,7,9,10}, 3
	res := missingElement(nums, k)
	fmt.Println(res)
	fmt.Println(2^2^2^3)
}
