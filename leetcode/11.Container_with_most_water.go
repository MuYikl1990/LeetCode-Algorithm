package main

import "fmt"

func maxArea(height []int) int {
	if len(height) == 0 {
		return 0
	}
	var res int
	left, right := 0, len(height)-1
	for left < right {
		if height[left] >= height[right] {
			res = Max(res, height[right] * (right - left))
			right--
		} else {
			res = Max(res, height[left] * (right - left))
			left++
		}
	}
	return res
}

func main() {
	height := []int{1,8,6,2,5,4,8,3,7}
	res := maxArea(height)
	fmt.Println(res)
}
