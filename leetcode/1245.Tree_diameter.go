package main

import "fmt"

func treeDiameter(edges [][]int) int {
	graph := make([][]int, len(edges) + 1)
	for _, edge := range edges {
		graph[edge[0]] = append(graph[edge[0]], edge[1])
		graph[edge[1]] = append(graph[edge[1]], edge[0])
	}

	res := 0
	var dfs1245 func([][]int, int, int) int
	dfs1245 = func(graph [][]int, pre, cur int) int {
		list := graph[cur]
		length := 0
		for j := 0; j < len(list); j++ {
			if pre != list[j] {
				length = max1245(length, dfs1245(graph, cur, list[j]))
			}
		}
		return length + 1
	}

	for i := 0; i < len(edges); i++ {
		if len(graph[i]) == 1 {
			res = max1245(res, dfs1245(graph, -1, i))
		}
	}

	return res - 1
}

func max1245(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	//edges := [][]int{{0,1}, {1,2}, {2,3}, {1,4}, {4,5}}
	edges := [][]int{{0,1}, {0,2}}
	res := treeDiameter(edges)
	fmt.Println(res)
}
