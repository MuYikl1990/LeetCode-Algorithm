package main

import "fmt"

func search(nums []int, target int) int {
	if len(nums) < 1 {
		return -1
	}

	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right - left) / 2
		if nums[mid] == target {
			return mid
		}

		if nums[left] <= nums[mid] {
			if nums[left] <= target && target < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			if nums[mid] < target && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	return -1
}

func main() {
	nums1, target1 := []int{4,5,6,7,0,1,2}, 2
	nums2, target2 := []int{4,5,6,7,0,1,2}, 5
	res1, res2 := search(nums1, target1), search(nums2, target2)
	fmt.Println(res1, res2)
}
