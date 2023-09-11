package main

import "fmt"

func cutWoods(nums []int, k int) int {
	left, right := 1, 0
	for i := range nums {
		if nums[i] > right {
			right = nums[i]
		}
	}

	//var canCut func(int) bool
	canCut := func(m int) bool {
		count := 0
		for _, num := range nums {
			count += num / m
		}
		return count >= k
	}

	for left < right {
		mid := left + (right - left + 1) / 2
		if canCut(mid) {
			left = mid
		} else {
			right = mid - 1
		}
	}
	return left
}

func main() {
	nums, k := []int{4,7,2,10,5}, 5
	res := cutWoods(nums, k)
	fmt.Println(res)
}
