package main

import "fmt"

func minPathSum(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	for i := 0; i < m; i++ {
		if i != 0 {
			grid[i][0] += grid[i-1][0]
		}
	}

	for j := 0; j < n; j++ {
		if j != 0 {
			grid[0][j] += grid[0][j-1]
		}
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			grid[i][j] = min64(grid[i-1][j], grid[i][j-1]) + grid[i][j]
		}
	}
	return grid[m-1][n-1]
}

func min64(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	grid := [][]int{{1,3,1}, {1,5,1}, {4,2,1}}
	res := minPathSum(grid)
	fmt.Println(res)
}
