package main

import "fmt"

func productExceptSelf(nums []int) []int {
	res := make([]int, len(nums))
	res[0] = 1
	for i := 1; i < len(nums); i++ {
		res[i] = res[i - 1] * nums[i - 1]
	}

	sum := 1
	for j := len(nums) - 1; j >= 0; j-- {
		res[j] *= sum
		sum *= nums[j]
	}
	return res
}

func main() {
	nums := []int{-1,1,0,-3,3}
	res := productExceptSelf(nums)
	fmt.Println(res)
}
