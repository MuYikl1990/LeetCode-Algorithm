package main

import "fmt"

func numDistinctIslands(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	row, col := len(grid), len(grid[0])
	islands := make(map[string]int)
	count := 0
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if grid[i][j] == 1 {
				var island string
				dfs694(grid, i, j, &island)
				if island != "" {
					islands[island]++
				}
			}
		}
	}

	for _ = range islands {
		count++
	}
	return count
}

func dfs694(grid [][]int, i, j int, island *string) {
	if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) || grid[i][j] == 0 {
		return
	}

	grid[i][j] = 0
	*island += "d"
	dfs694(grid, i + 1, j, island)
	*island += "u"
	dfs694(grid, i - 1, j, island)
	*island += "r"
	dfs694(grid, i, j + 1, island)
	*island += "l"
	dfs694(grid, i, j - 1, island)
}

func main() {
	grid := [][]int{
		{1, 1, 0, 1, 1},
		{1, 0, 0, 0, 0},
		{1, 1, 0, 0, 1},
		{0, 1, 0, 1, 1},
		{1, 0, 0, 0, 0},
	}
	res := numDistinctIslands(grid)
	fmt.Println(res)
}
