package main

import "fmt"

func largestRectangleArea(heights []int) int {
	size := len(heights)
	heights = append(heights, 0)
	heights = append([]int{0}, heights...)
	res := 0
	stack := []int{0}

	for i := 1; i < size + 2; i++ {
		for heights[i] < heights[stack[len(stack) - 1]] {
			height := heights[stack[len(stack) - 1]]
			stack = stack[:len(stack) - 1]
			width := i - stack[len(stack) - 1] - 1
			res = max84(res, height * width)
		}
		stack = append(stack, i)
	}
	return res
}

func max84(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	heights := []int{2,1,5,6,2,3}
	res := largestRectangleArea(heights)
	fmt.Println(res)
}
