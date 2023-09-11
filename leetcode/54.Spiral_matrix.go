package main

import "fmt"

func spiralOrder(matrix [][]int) []int {
	var res []int
	if len(matrix) == 0 {
		return res
	}
	m, n := len(matrix), len(matrix[0])
	up, down, left, right := 0, m - 1, 0, n - 1
	for {
		for i := left; i <= right; i++ {
			res = append(res, matrix[up][i])
		}
		up++
		if up > down {
			break
		}

		for i := up; i <= down; i++ {
			res = append(res, matrix[i][right])
		}
		right--
		if left > right {
			break
		}

		for i := right; i >= left; i-- {
			res = append(res, matrix[down][i])
		}
		down--
		if up > down {
			break
		}

		for i := down; i >= up; i-- {
			res = append(res, matrix[i][left])
		}
		left++
		if left > right {
			break
		}
	}
	return res
}

func main() {
	matrix := [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}
	res := spiralOrder(matrix)
	fmt.Println(res)
}
