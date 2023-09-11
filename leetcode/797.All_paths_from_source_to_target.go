package main

import "fmt"

func allPathsSourceTarget(graph [][]int) [][]int {
	n := len(graph)
	res := make([][]int, 0)
	var dfs func(int, []int)
	dfs = func(index int, tmp []int) {
		if index == n - 1 {
			ans := make([]int, len(tmp))
			copy(ans, tmp)
			res = append(res, ans)
			return
		}
		for i := range graph[index] {
			tmp = append(tmp, graph[index][i])
			dfs(graph[index][i], tmp)
			tmp = tmp[:len(tmp) - 1]
		}
	}
	dfs(0, []int{0})
	return res
}

func main() {
	graph := [][]int{{4,3,1}, {3,2,4}, {3}, {4}, {}}
	res := allPathsSourceTarget(graph)
	fmt.Println(res)
}
