package main

import "fmt"

func searchMatrix(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])
	left, right := 0, n - 1
	for left < m && right >= 0 {
		if matrix[left][right] > target {
			right--
		} else if matrix[left][right] < target {
			left++
		} else {
			return true
		}
	}
	return false
}

// dichotomy
func searchMatrixByDichotomy(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])
	left, right := 0, m * n - 1
	for left < right {
		mid := left + (right - left) / 2
		num := matrix[mid / n][mid % n]
		if num < target {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return matrix[left / n][left % n] == target
}

func main() {
	matrix, target := [][]int{{1, 3, 5, 7},{10, 11, 16, 20},{23, 30, 34, 60}}, 3
	res := searchMatrix(matrix, target)
	fmt.Println(res)
}
