package main

import "fmt"

func findNumberIn2DArray(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	n := len(matrix)
	m := len(matrix[0])

	for i, j := 0, m - 1; j >= 0 && i < n; {
		if target > matrix[i][j] {
			i++
		} else if target < matrix[i][j] {
			j--
		} else {
			return true
		}
	}
	return false
}

func main() {
	matrix := [][]int{{1, 4, 7, 11, 15}, {2, 5, 8, 12, 19}, {3, 6, 9, 16, 22}, {10, 13, 14, 17, 24}, {18, 21, 23, 26, 30}}
	target := 20
	res := findNumberIn2DArray(matrix, target)
	fmt.Println(res)
}
