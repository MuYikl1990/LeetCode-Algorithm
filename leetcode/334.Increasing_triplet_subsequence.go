package main

import (
	"fmt"
	"math"
)

func increasingTriplet(nums []int) bool {
	if len(nums) < 3 {
		return false
	}
	fir, sec := math.MaxInt, math.MaxInt

	for i := range nums {
		if nums[i] <= fir {
			fir = nums[i]
		} else if nums[i] <= sec {
			sec = nums[i]
		} else if nums[i] > sec {
			return true
		}
	}
	return false
}

func main() {
	nums := []int{20,100,10,12,5,13}
	res := increasingTriplet(nums)
	fmt.Println(res)
}
