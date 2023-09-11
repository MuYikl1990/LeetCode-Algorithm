package main

import "fmt"

func numIslands(grid [][]byte) int {
	res := 0
	for r := 0; r < len(grid); r++ {
		for c := 0; c < len(grid[r]); c++ {
			if grid[r][c] == '1' {
				dfs200(grid, r, c)
				res++
			}
		}
	}
	return res
}

func dfs200(grid [][]byte, r, c int) {
	if r < 0 || r >= len(grid) || c < 0 || c >= len(grid[r]) {
		return
	}

	if grid[r][c] != '1' {
		return
	}
	grid[r][c] = '2'
	dfs200(grid, r, c-1)
	dfs200(grid, r, c+1)
	dfs200(grid, r-1, c)
	dfs200(grid, r+1, c)
}

func main() {
	grid := [][]byte{{'1','1','0','0','0'}, {'1','1','0','0','0'}, {'0','0','1','0','0'}, {'0','0','0','1','1'}}
	res := numIslands(grid)
	fmt.Println(res)
}
