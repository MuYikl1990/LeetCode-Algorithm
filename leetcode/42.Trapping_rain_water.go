package main

import "fmt"

func trap(height []int) int {
	left, right := 0, len(height) - 1
	leftMax, rightMax := 0, 0
	res := 0

	for left <= right {
		if leftMax < rightMax {
			res += max42(0, leftMax - height[left])
			leftMax = max42(leftMax, height[left])
			left++
		} else {
			res += max42(0, rightMax - height[right])
			rightMax = max42(rightMax, height[right])
			right--
		}
	}
	return res
}

func max42(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	height := []int{4,2,0,3,2,5}
	res := trap(height)
	fmt.Println(res)
}
