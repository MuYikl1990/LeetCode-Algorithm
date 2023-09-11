package main

import "fmt"

func searchMatrixII(matrix [][]int, target int) bool {
	i, j := 0, len(matrix[0]) - 1
	for i < len(matrix) && j >= 0 {
		if target == matrix[i][j] {
			return true
		} else if target > matrix[i][j] {
			i++
		} else {
			j--
		}
	}
	return false
}

func main() {
	matrix, target := [][]int{{1,4,7,11,15}, {2,5,8,12,19}, {3,6,9,16,22}, {10,13,14,17,24}, {18,21,23,26,30}}, 20
	res := searchMatrixII(matrix, target)
	fmt.Println(res)
}
