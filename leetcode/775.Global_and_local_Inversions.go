package main

import (
	"fmt"
	"math"
)

func isIdealPermutation(nums []int) bool {
	for i := range nums {
		if math.Abs(float64(nums[i] - i)) > 1 {
			return false
		}
	}
	return true
}

func main() {
	nums := []int{7,3,1,0,2,6,5,4}
	res := isIdealPermutation(nums)
	fmt.Println(res)
}
