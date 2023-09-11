package main

import (
	"fmt"
	"reflect"
)

func setZeroes(matrix [][]int) {
	m, n := len(matrix), len(matrix[0])
	rFlag, cFlag := false, false
	for i := 0; i < m; i++ {
		if matrix[i][0] == 0 {
			cFlag = true
			break
		}
	}
	for j := 0; j < n; j++ {
		if matrix[0][j] == 0 {
			rFlag = true
			break
		}
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[i][j] == 0 {
				matrix[i][0], matrix[0][j] = 0, 0
			}
		}
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
	}
	if rFlag {
		for j := 0; j < n; j++ {
			matrix[0][j] = 0
		}
	}
	if cFlag {
		for i := 0; i < m; i++ {
			matrix[i][0] = 0
		}
	}
}

func main() {
	matrix := [][]int{{0,1,2,0}, {3,4,5,2}, {1,3,1,5}}
	setZeroes(matrix)
	expected := [][]int{{0,0,0,0}, {0,4,5,0},{0,3,1,0}}
	fmt.Println(reflect.DeepEqual(matrix, expected))
}
