package main

import "fmt"

// 现有n个编译项，编号为0 ~ n-1。给定一个二维数组，表示编译项之间有依赖关系。如[0, 1]表示1依赖于0
// 若存在循环依赖则返回空；不存在依赖则返回可行的编译顺序。
func haveCircularDependency(n int, prerequisites [][]int) []int {
	graph := make(map[int][]int)
	indegree := make([]int, n)
	var res []int
	for _, prerequisite := range prerequisites {
		graph[prerequisite[0]] = append(graph[prerequisite[0]], prerequisite[1])
		indegree[prerequisite[1]]++
	}

	var queue []int
	for i := range indegree {
		if indegree[i] == 0 {
			queue = append(queue, i)
		}
	}

	for len(queue) > 0 {
		cur := queue[0]
		res = append(res, cur)
		for i := 0; i < len(graph[cur]); i++ {
			indegree[graph[cur][i]]--
			if indegree[graph[cur][i]] == 0 {
				queue = append(queue, graph[cur][i])
			}
		}
		queue = queue[1:]
	}
	if len(res) == n {
		return res
	}
	return []int{}
}

func main() {
	n, prerequisites := 5, [][]int{{0,2}, {1,2}, {2,3}, {2,1}}
	res := haveCircularDependency(n, prerequisites)
	fmt.Println(res)
}
