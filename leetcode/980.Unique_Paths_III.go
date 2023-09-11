package main

import "fmt"

func uniquePathsIII(grid [][]int) int {
	tot, m, n, res := 1, len(grid), len(grid[0]), 0
	startX, startY := 0, 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				startX, startY = i, j
				continue
			}
			if grid[i][j] == 0 {
				tot++
			}
		}
	}
	backtrace980(grid, startX, startY, tot, &res)
	return res
}

func backtrace980(grid [][]int, i, j, tot int, res *int) {
	if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) || grid[i][j] == -1 {
		return
	}
	if grid[i][j] == 2 {
		if tot == 0 {
			*res++
			return
		} else {
			return
		}
	}

	grid[i][j] = -1
	dir := []int{-1, 0, 1, 0, -1}
	for k := 0; k < 4; k++ {
		newX, newY := i + dir[k], j + dir[k + 1]
		backtrace980(grid, newX, newY, tot - 1, res)
	}
	grid[i][j] = 0
}

func main() {
	grid := [][]int{{1,0,0,0}, {0,0,0,0}, {0,0,0,2}}
	res := uniquePathsIII(grid)
	fmt.Println(res)
}
