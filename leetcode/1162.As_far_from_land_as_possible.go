package main

import "fmt"

func maxDistance(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	var queue []ele
	tmp, flag := ele{-1, -1}, false
	dir := [5]int{-1, 0, 1, 0, -1}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				queue = append(queue, ele{i, j})
			}
		}
	}

	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		tmp.x, tmp.y = cur.x, cur.y
		for i := 0; i < 4; i++ {
			nx := tmp.x + dir[i]
			ny := tmp.y + dir[i + 1]
			if nx < 0 || ny < 0 || nx >= n || ny >= m || grid[nx][ny] != 0 {
				continue
			}
			flag = true
			queue = append(queue, ele{nx, ny})
			grid[nx][ny] = grid[tmp.x][tmp.y] + 1
		}
	}

	if tmp.x == -1 || !flag {
		return -1
	}
	return grid[tmp.x][tmp.y] - 1
}

type ele struct {
	x int
	y int
}

func main() {
	grid := [][]int{{1,0,1}, {0,0,0}, {1,0,1}}
	res := maxDistance(grid)
	fmt.Println(res)
}
