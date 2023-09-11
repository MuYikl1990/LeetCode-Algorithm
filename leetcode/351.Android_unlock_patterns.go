package main

import "fmt"

func numberOfPatterns(m, n int) int {
	visited := make([]bool, 10)
	middle := make([][]int, 10)
	for i := range middle {
		middle[i] = make([]int, 10)
	}
	middle[1][3], middle[3][1] = 2, 2
	middle[1][7], middle[7][1] = 4, 4
	middle[3][9], middle[9][3] = 6, 6
	middle[7][9], middle[9][7] = 8, 8
	middle[2][8], middle[8][2], middle[4][6], middle[6][4] = 5, 5, 5, 5
	res := 0
	for i := m; i <= n; i++ {
		res += dfs351(1, m - 1, visited, middle) * 4
		res += dfs351(2, m - 1, visited, middle) * 4
		res += dfs351(5, m - 1, visited, middle)
	}
	return res
}

func dfs351(current, keys int, visited []bool, middle [][]int) int {
	if keys == 0 {
		return 1
	}

	result := 0
	visited[current] = true
	for next := 1; next <= 9; next++ {
		jump := middle[current][next]
		if !visited[next] && (jump == 0 || visited[jump]) {
			result += dfs351(next, keys - 1, visited, middle)
		}
	}
	visited[current] = false
	return result
}

func main() {
	m, n := 1, 3
	res := numberOfPatterns(m , n)
	fmt.Println(res)
}
