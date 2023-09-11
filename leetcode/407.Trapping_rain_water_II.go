package main

import (
	"container/heap"
	"fmt"
)

func trapRainWater(heightMap [][]int) int {
	m, n := len(heightMap), len(heightMap[0])

	if m < 3 || n < 3 {
		return 0
	}

	minQ := &minHeap{}
	res, visited := 0, make([][]bool, m)
	for k := range visited {
		visited[k] = make([]bool, n)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 || i == m - 1 || j == 0 || j == n - 1 {
				heap.Push(minQ, cell{i, j, heightMap[i][j]})
				visited[i][j] = true
			}
		}
	}

	dir := [5]int{-1, 0, 1, 0, -1}
	for minQ.Len() > 0 {
		water := heap.Pop(minQ)
		for i := 0; i < 4; i++ {
			nx := water.(cell).x + dir[i]
			ny := water.(cell).y + dir[i + 1]
			if nx >= 0 && nx <= m - 1 && ny >= 0 && ny <= n - 1 && visited[nx][ny] != true {
				if heightMap[nx][ny] < water.(cell).h {
					res += water.(cell).h - heightMap[nx][ny]
				}
				heap.Push(minQ, cell{nx, ny, max407(heightMap[nx][ny], water.(cell).h)})
				visited[nx][ny] = true
			}
		}
	}
	return res
}

func max407(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type cell struct {
	x, y, h int
}

type minHeap []cell

func (hp *minHeap) Len() int {
	return len(*hp)
}

func (hp *minHeap) Less(i, j int) bool {
	return (*hp)[i].h < (*hp)[j].h
}

func (hp *minHeap) Swap(i, j int) {
	(*hp)[i], (*hp)[j] = (*hp)[j], (*hp)[i]
}

func (hp *minHeap) Push(x interface{}) {
	*hp = append(*hp, x.(cell))
}

func (hp *minHeap) Pop() interface{} {
	num := (*hp)[len(*hp) - 1]
	*hp = (*hp)[:len(*hp) - 1]
	return num
}

func main() {
	heightMap := [][]int{{1,4,3,1,3,2}, {3,2,1,3,2,4}, {2,3,3,2,3,1}}
	res := trapRainWater(heightMap)
	fmt.Println(res)
}
