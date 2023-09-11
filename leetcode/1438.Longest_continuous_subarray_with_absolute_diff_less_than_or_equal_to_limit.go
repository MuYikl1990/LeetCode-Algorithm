package main

import (
	"fmt"
)

func longestSubarray(nums []int, limit int) int {
	// monotonic stack
	minQ := make([]int, 0)
	maxQ := make([]int, 0)
	left, right := 0, 0
	res := 0
	for ;right < len(nums); right++ {
		for len(minQ) > 0 && minQ[len(minQ) - 1] > nums[right] {
			minQ = minQ[:len(minQ) - 1]
		}
		minQ = append(minQ, nums[right])
		for len(maxQ) > 0 && maxQ[len(maxQ) - 1] < nums[right] {
			maxQ = maxQ[:len(maxQ) - 1]
		}
		maxQ = append(maxQ, nums[right])
		for len(minQ) > 0 && len(maxQ) > 0 && maxQ[0] - minQ[0] > limit {
			if nums[left] == minQ[0] {
				minQ = minQ[1:]
			}
			if nums[left] == maxQ[0] {
				maxQ = maxQ[1:]
			}
			left++
		}
		res = max1438(res, right - left + 1)
	}
	return res
}

func max1438(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	nums, limit := []int{10,1,2,4,7,2}, 5
	res := longestSubarray(nums, limit)
	fmt.Println(res)
	fmt.Println(6&0)
}
