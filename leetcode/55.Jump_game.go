package main

import "fmt"

func canJump(nums []int) bool {
	length := 0
	for i := range nums {
		if i > length {
			return false
		}
		length = max55(length, i + nums[i])
	}
	return true
}

func max55(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	nums := []int{2,3,1,1,4}
	res := canJump(nums)
	fmt.Println(res)
}
