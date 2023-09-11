package main

import "fmt"

func orangesRotting(grid [][]int) int {
	count := 0
	var queue [][]int
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 1 {
				count++
			} else if grid[i][j] == 2 {
				queue = append(queue, []int{i, j})
			}
		}
	}
	if count == 0 {
		return 0
	}

	dirs := [5]int{-1, 0, 1, 0, -1}
	minute := 0
	for len(queue) > 0 {
		n := len(queue)
		minute++
		for i := 0; i < n; i++ {
			for k := 0; k < 4; k++ {
				row, col := queue[i][0] + dirs[k], queue[i][1] + dirs[k + 1]
				if row >= 0 && col >= 0 && row < len(grid) && col < len(grid[0]) && grid[row][col] == 1 {
					count--
					grid[row][col] = 2
					queue = append(queue, []int{row, col})
				}
			}
		}
		queue = queue[n:]
	}
	if count > 0 {
		return -1
	}
	return minute - 1
}

func main() {
	grid := [][]int{{2,1,1}, {1,1,1}, {0,1,2}}
	res := orangesRotting(grid)
	fmt.Println(res)
}
