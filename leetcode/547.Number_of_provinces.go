package main

import "fmt"

func add(x int, nodes map[int]int) {
	if _, ok := nodes[x]; !ok {
		nodes[x] = x
	}
}

func findNode(x int, nodes map[int]int) int {
	root := x
	for root != nodes[root] {
		root = nodes[root]
	}

	for x != root {
		tmp := nodes[x]
		nodes[x] = root
		x = tmp
	}
	return root
}

func union(x, y int, nodes map[int]int) {
	xRoot, yRoot := findNode(x, nodes), findNode(y, nodes)
	if xRoot != yRoot {
		nodes[xRoot] = yRoot
	}
}

func findCircleNum(isConnected [][]int) int {
	nodes, count := make(map[int]int), 0
	for i := 0; i < len(isConnected); i++ {
		add(i, nodes)
		for j := 0; j < i; j++ {
			if isConnected[i][j] == 1 {
				union(i, j, nodes)
			}
		}
	}
	for key, value := range nodes {
		if value == key {
			count++
		}
	}
	return count
}

func main() {
	isConnected := [][]int{{1,0,0,1}, {0,1,1,0}, {0,1,1,1}, {1,0,1,1}}
	isConnected1 := [][]int{{1,1,0}, {1,1,0}, {0,0,1}}
	res := findCircleNum(isConnected)
	res1 := findCircleNum(isConnected1)
	fmt.Println(res)
	fmt.Println(res1)
}
