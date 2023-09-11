package main

import "fmt"

type NumMatrix struct {
	preSum [][]int
}

func Constructor304(matrix [][]int) NumMatrix {
	m, n := len(matrix), len(matrix[0])
	preSum := make([][]int, m + 1)
	for k := range preSum {
		preSum[k] = make([]int, n + 1)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			preSum[i + 1][j + 1] = preSum[i + 1][j] + preSum[i][j + 1] - preSum[i][j] + matrix[i][j]
		}
	}
	return NumMatrix{preSum}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	return this.preSum[row2 + 1][col2 + 1] - this.preSum[row1][col2 + 1] - this.preSum[row2 + 1][col1] + this.preSum[row1][col1]
}

func main() {
	matrix := [][]int{{3, 0, 1, 4, 2}, {5, 6, 3, 2, 1}, {1, 2, 0, 1, 5}, {4, 1, 0, 1, 7}, {1, 0, 3, 0, 5}}
	obj := Constructor304(matrix)
	row1,col1,row2,col2 := 2, 1, 4, 3
	param1 := obj.SumRegion(row1,col1,row2,col2)
	fmt.Println(param1)
	row1,col1,row2,col2 = 1, 1, 2, 2
	param2 := obj.SumRegion(row1,col1,row2,col2)
	fmt.Println(param2)
	row1,col1,row2,col2 = 1, 2, 2, 4
	param3 := obj.SumRegion(row1,col1,row2,col2)
	fmt.Println(param3)
}
