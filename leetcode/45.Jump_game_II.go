package main

import "fmt"

func jump(nums []int) int {
	maxPos, end, res := 0, 0, 0
	for i := 0; i < len(nums) - 1; i++ {
		maxPos = max45(maxPos, i + nums[i])
		if i == end {
			end = maxPos
			res++
		}
	}
	return res
}

func max45(a,b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	nums := []int{2,3,1,1,4}
	res := jump(nums)
	fmt.Println(res)
}
