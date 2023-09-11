package main

import "fmt"

func maximalNetworkRank(n int, roads [][]int) int {
	rank := make(map[int]map[int]int)
	inDegree := make([]int, n)
	res := 0

	for _, road := range roads {
		if rank[road[0]] == nil {
			rank[road[0]] = make(map[int]int)
		}
		rank[road[0]][road[1]]++

		if rank[road[1]] == nil {
			rank[road[1]] = make(map[int]int)
		}
		rank[road[1]][road[0]]++

		inDegree[road[0]]++
		inDegree[road[1]]++
	}

	for i := 0; i < n - 1; i++ {
		for j := i + 1; j < n; j++ {
			res = max1615(res, inDegree[i] + inDegree[j] - rank[i][j])
		}
	}
	return res
}

func max1615(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	n, roads := 4, [][]int{{0, 1}, {0, 3}, {1, 2}, {1, 3}}
	res := maximalNetworkRank(n, roads)
	fmt.Println(res)
}
