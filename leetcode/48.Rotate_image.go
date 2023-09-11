package main

import (
	"fmt"
	"reflect"
)

func rotate(matrix [][]int) [][]int {
	n := len(matrix)
	for i := 0; i < n / 2; i++ {
		matrix[i], matrix[n - i - 1] = matrix[n - i - 1], matrix[i]
	}
	for j := 0; j < n; j++ {
		for k := j; k < n; k++ {
			matrix[j][k], matrix[k][j] = matrix[k][j], matrix[j][k]
		}
	}
	return matrix
}

func main() {
	matrix := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	res := rotate(matrix)
	value := [][]int{{7, 4, 1}, {8, 5, 2}, {9, 6, 3}}
	fmt.Println(reflect.DeepEqual(res, value))
}
