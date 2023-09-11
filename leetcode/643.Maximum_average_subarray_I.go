package main

import (
	"fmt"
	"math"
)

// preSum
func findMaxAverage(nums []int, k int) float64 {
	res := math.MinInt
	n := len(nums)
	preSum := make([]int, n + 1)
	for i := range nums {
		preSum[i + 1] = nums[i] + preSum[i]
	}

	for i := n; i >= k; i-- {
		res = max643(res, preSum[i] - preSum[i - k])
	}
	return float64(res) / float64(k)
}

// Slide Window
func findMaxAverage1(nums []int, k int) float64 {
	left, right, sum, res := 0, 0, 0, math.MinInt
	for right < len(nums) {
		sum += nums[right]
		if right - left == k - 1 {
			res = max643(res, sum)
			sum -= nums[left]
			left++
		}
		right++
	}
	return float64(res) / float64(k)
}


func max643(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	nums, k := []int{1,12,-5,-6,50,3}, 4
	res := findMaxAverage1(nums, k)
	fmt.Println(res)
}
