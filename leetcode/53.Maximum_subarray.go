package main

import (
	"fmt"
	"math"
)

func maxSubArray(nums []int) int {
	sum := math.MinInt8
	res := sum
	for _, value := range nums {
		sum = max53(sum + value, value)
		res = max53(res, sum)
	}
	return res
}

func max53(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	nums := []int{-2,1,-3,4,-1,2,1,-5,4}
	res := maxSubArray(nums)
	fmt.Println(res)
}
