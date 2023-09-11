package main

import "fmt"

func maximalRectangle(matrix [][]byte) int {
	if len(matrix) == 0 {
		return 0
	}

	heights := make([]int, len(matrix[0]) + 1)
	res := 0
	for row := 0; row < len(matrix); row++ {
		var stack []int
		for col := 0; col <= len(matrix[0]); col++ {
			if col < len(matrix[0]) {
				if matrix[row][col] == '1' {
					heights[col]++
				} else {
					heights[col] = 0
				}
			}
			for len(stack) != 0 && heights[col] < heights[stack[len(stack) - 1]] {
				height := heights[stack[len(stack) - 1]]
				stack = stack[:len(stack) - 1]
				diff := -1
				if len(stack) != 0 {
					diff = stack[len(stack) - 1]
				}
				width := col - diff - 1
				res = max85(res, height * width)
			}
			stack = append(stack, col)
		}
	}
	return res
}

func max85(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	matrix := [][]byte{{'1','0','1','0','0'}, {'1','0','1','1','1'}, {'1','1','1','1','1'}, {'1','0','0','1','0'}}
	res := maximalRectangle(matrix)
	fmt.Println(res)
}
