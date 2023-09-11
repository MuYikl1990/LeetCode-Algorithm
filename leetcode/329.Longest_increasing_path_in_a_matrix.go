package main

import "fmt"

func longestIncreasingPath(matrix [][]int) int {
	if len(matrix) == 0 {
		return 0
	}
	m, n, res := len(matrix), len(matrix[0]), 0
	memo := make([][]int, m)
	for i := range memo {
		memo[i] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if memo[i][j] == 0 {
				res = max329(res, dfs329(i, j, m, n, matrix, memo))
			}
		}
	}
	return res
}

func dfs329(i, j, m, n int, matrix, memo [][]int) int {
	dir := [5]int{-1, 0, 1, 0, -1}
	if memo[i][j] != 0 {
		return memo[i][j]
	}

	ans := 1
	for k := 0; k < 4; k++ {
		x := i + dir[k]
		y := j + dir[k+1]
		if x >= 0 && x < m && y >= 0 && y < n && matrix[i][j] < matrix[x][y] {
			ans = max329(ans, dfs329(x, y, m, n, matrix, memo) + 1)
		}
	}
	memo[i][j] = ans
	return ans
}

func max329(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	matrix := [][]int{{9,9,4}, {6,6,8}, {2,1,1}}
	res := longestIncreasingPath(matrix)
	fmt.Println(res)
}
