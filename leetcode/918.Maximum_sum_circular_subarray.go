package main

import "fmt"

func maxSubarraySumCircular(nums []int) int {
	tot, curMax, maxSum, curMin, minSum := 0, 0, nums[0], 0, nums[0]
	for i := range nums {
		curMax = max918(curMax + nums[i], nums[i])
		maxSum = max918(maxSum, curMax)
		curMin = min918(curMin + nums[i], nums[i])
		minSum = min918(minSum, curMin)
		tot += nums[i]
	}
	if maxSum < 0 {
		return maxSum
	}
	return max918(maxSum, tot - minSum)
}

func min918(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max918(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	nums := []int{5, -3, 5}
	res := maxSubarraySumCircular(nums)
	fmt.Println(res)
}
