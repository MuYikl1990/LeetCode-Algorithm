package main

import "fmt"

func maximalSquare(matrix [][]byte) int {
	res := 0
	m, n := len(matrix), len(matrix[0])
	dp := make([][]int, m + 1)
	for i := range dp {
		dp[i] = make([]int, n + 1)
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if matrix[i - 1][j - 1] == '1' {
				dp[i][j] = min221(dp[i - 1][j], min221(dp[i][j - 1], dp[i - 1][j - 1])) + 1
			}
			res = max221(res, dp[i][j])
		}
	}

	return res * res
}

func max221(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min221(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	matrix := [][]byte{{'1','0','1','0','0'}, {'1','0','1','1','1'}, {'1','1','1','1','1'}, {'1','0','0','1','0'}}
	res := maximalSquare(matrix)
	fmt.Println(res)
}
