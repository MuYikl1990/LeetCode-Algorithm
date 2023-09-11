package main

import "fmt"

func removeElement(nums []int, val int) int {
	left, right := 0, len(nums) - 1
	for left <= right {
		if nums[left] != val {
			left++
		} else {
			if nums[right] != val {
				nums[left] = nums[right]
				left++
			}
			right--
		}

	}
	return left
}

func main() {
	nums, val := []int{1}, 1
	res := removeElement(nums, val)
	fmt.Println(res)
}
