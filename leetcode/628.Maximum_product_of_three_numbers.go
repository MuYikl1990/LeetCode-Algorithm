package main

import (
	"fmt"
	"math"
)

func maximumProduct(nums []int) int {
	max1, max2, max3, min1, min2 := math.MinInt, math.MinInt, math.MinInt, math.MaxInt, math.MaxInt
	for _, num := range nums {
		if num > max1 {
			max3 = max2
			max2 = max1
			max1 = num
		} else if num > max2 {
			max3 = max2
			max2 = num
		} else if num > max3 {
			max3 = num
		}

		if num < min1 {
			min2 = min1
			min1 = num
		} else if num < min2 {
			min2 = num
		}
	}
	return max628(max1 * max2 * max3, max1 * min1 * min2)
}

func max628(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	nums := []int{1,2,3,4}
	res := maximumProduct(nums)
	fmt.Println(res)
}
