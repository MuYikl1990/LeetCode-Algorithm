package main

import (
	"fmt"
	"math"
)

// Dijkstra 算法
func networkDelayTime(times [][]int, n int, k int) int {
	graph := make([][]int, n)
	for i := range graph {
		graph[i] = make([]int, n)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			graph[i][j] = math.MaxInt / 2
		}
	}
	for _, time := range times {
		graph[time[0] - 1][time[1] - 1] = time[2]
	}

	distant := make([]int, n)
	for i := range distant {
		distant[i] = math.MaxInt / 2
	}
	distant[k - 1] = 0

	used := make([]bool, n)

	for i := 0; i < n; i++ {
		x := -1
		for y := 0; y < n; y++ {
			if !used[y] && (x == -1 || distant[y] < distant[x]) {
				x = y
			}
		}

		used[x] = true
		for z := 0; z < n; z++ {
			distant[z] = min743(distant[z], distant[x] + graph[x][z])
		}
	}
	res := distant[0]
	for i := range distant {
		res = max743(res, distant[i])
	}
	if res == math.MaxInt / 2 {
		return -1
	}
	return res
}

func min743(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max743(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	times, n, k := [][]int{{2,1,1}, {2,3,1}, {3,4,1}}, 4, 2
	res := networkDelayTime(times, n, k)
	fmt.Println(res)
}
