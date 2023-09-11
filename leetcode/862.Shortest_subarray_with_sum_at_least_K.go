package main

import (
	"fmt"
	"math"
)

func shortestSubarray(nums []int, k int) int {
	preSum := make([]int, len(nums) + 1)
	for i := range nums {
		preSum[i + 1] = preSum[i] + nums[i]
	}

	res := math.MaxInt
	var queue []int
	for i, sum := range preSum {
		for len(queue) > 0 && sum - preSum[queue[0]] >= k {
			res = min862(res, i - queue[0])
			queue = queue[1:]
		}
		for len(queue) > 0 && sum <= preSum[queue[len(queue) - 1]] {
			queue = queue[:len(queue) - 1]
		}
		queue = append(queue, i)
	}
	if res < math.MaxInt {
		return res
	}
	return -1
}

func min862(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	nums, k := []int{2,-1,2}, 3
	res := shortestSubarray(nums, k)
	fmt.Println(res)
}
