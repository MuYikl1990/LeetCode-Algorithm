package main

import "fmt"

func maxAreaOfIsland(grid [][]int) int {
	res := 0
	var dfs695 func([][]int, int, int) int
	dfs695 = func(grid [][]int, row, col int) int {
		if row >= len(grid) || row < 0 || col >= len(grid[row]) || col < 0 || grid[row][col] != 1 {
			return 0
		}
		grid[row][col] = 2
		dirs := []int{-1, 0, 1, 0, -1}
		area := 1
		for k := 0; k < 4; k++ {
			newRow := row + dirs[k]
			newCol := col + dirs[k + 1]
			area += dfs695(grid, newRow, newCol)
		}
		return area
	}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 1 {
				res = max695(res, dfs695(grid, i, j))
			}
		}
	}
	return res
}

func max695(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	grid := [][]int{{0,0,1,0,0,0,0,1,0,0,0,0,0},{0,0,0,0,0,0,0,1,1,1,0,0,0},{0,1,1,0,1,0,0,0,0,0,0,0,0},
					{0,1,0,0,1,1,0,0,1,0,1,0,0},{0,1,0,0,1,1,0,0,1,1,1,0,0},{0,0,0,0,0,0,0,0,0,0,1,0,0},{
					0,0,0,0,0,0,0,1,1,1,0,0,0},{0,0,0,0,0,0,0,1,1,0,0,0,0}}
	res := maxAreaOfIsland(grid)
	fmt.Println(res)
}
