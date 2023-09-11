package main

import "fmt"

func nextGreaterElements(nums []int) []int {
	var stack []int
	n := len(nums)
	res := make([]int, n)
	for k := range res {
		res[k] = -1
	}

	for i := 0; i < 2 * n; i++ {
		for len(stack) > 0 && nums[stack[len(stack) - 1]] < nums[i % n] {
			res[stack[len(stack) - 1]] = nums[i % n]
			stack = stack[:len(stack) - 1]
		}
		stack = append(stack, i % n)
	}
	return res
}

func main() {
	nums := []int{1,2,3,4,3}
	res := nextGreaterElements(nums)
	fmt.Println(res)
}
