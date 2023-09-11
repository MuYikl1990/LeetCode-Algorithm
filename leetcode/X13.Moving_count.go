package main

import "fmt"

func movingCount(m int, n int, k int) int {
	visited := make([][]bool, m)
	for i := range visited {
		visited[i] = make([]bool, n)
	}

	var dfs func(int, int) int
	dfs = func(row, col int) int {
		if row < 0 || col < 0 || row >= m || col >= n || visited[row][col] || (row / 10 + row % 10 + col / 10 + col % 10) > k {
			return 0
		}
		visited[row][col] = true
		count := 1
		dirs := []int{-1, 0, 1, 0, -1}
		for i := 0; i < 4; i++ {
			count += dfs(row + dirs[i], col + dirs[i + 1])
		}
		return count
	}
	res := dfs(0, 0)
	return res
}

func main() {
	m, n, k := 3, 1, 0
	res := movingCount(m, n, k)
	fmt.Println(res)
}
