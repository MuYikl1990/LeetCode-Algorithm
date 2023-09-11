package main

import "fmt"

func matrixBlockSum(mat [][]int, k int) [][]int {
	m, n := len(mat), len(mat[0])
	res := make([][]int, m)
	for i := range res {
		res[i] = make([]int, n)
	}

	preSum := make([][]int, m + 1)
	for i := range preSum {
		preSum[i] = make([]int, n + 1)
	}
	for i := 1; i <= m; i ++ {
		for j := 1; j <= n; j++ {
			preSum[i][j] = preSum[i - 1][j] + preSum[i][j - 1] - preSum[i - 1][j - 1] + mat[i - 1][j - 1]
		}
	}

	for row := 0; row < m; row++ {
		for col := 0; col < n; col++ {
			row1, col1 := max1314(row - k, 0), max1314(col - k, 0)
			row2, col2 := min1314(row + k, m - 1), min1314(col + k, n - 1)
			res[row][col] = preSum[row2 + 1][col2 + 1] - preSum[row2 + 1][col1] - preSum[row1][col2 + 1] + preSum[row1][col1]
		}
	}
	return res
}

func max1314(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min1314(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	mat, k := [][]int{{1,2,3}, {4,5,6}, {7,8,9}}, 1
	res := matrixBlockSum(mat, k)
	fmt.Println(res)
}
