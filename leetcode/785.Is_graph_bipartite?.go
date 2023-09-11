package main

import "fmt"

func isBipartite(graph [][]int) bool {
	n := len(graph)
	visited := make([]int, n)

	for i := 0; i < n; i++ {
		if visited[i] != 0 {
			continue
		}

		stack := []int{i}
		visited[i] = 1

		for len(stack) > 0 {
			cur := stack[0]
			stack = stack[1:]
			flag := -1 * visited[cur]
			for _, ele := range graph[cur] {
				if visited[ele] == 0 {
					visited[ele] = flag
					stack = append(stack, ele)
				} else if visited[ele] != flag {
					return false
				}
			}
		}
	}
	return true
}

func main() {
	graphs := [][]int{{1,3}, {0,2}, {1,3}, {0,2}}
	res := isBipartite(graphs)
	fmt.Println(res)
}
