package main

import (
	"fmt"
	"math"
)

func maxProduct(nums []int) int {
	res, imax, imin := math.MinInt, 1, 1
	for _, val := range nums {
		if val < 0 {
			imax, imin = imin, imax
		}
		imax = max152(imax * val, val)
		imin = min152(imin * val, val)
		res = max152(res, imax)
	}
	return res
}

func max152(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min152(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	nums := []int{2,1,-2,4}
	res := maxProduct(nums)
	fmt.Println(res)
}
