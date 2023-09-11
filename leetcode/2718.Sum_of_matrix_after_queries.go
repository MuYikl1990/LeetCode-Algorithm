package main

import "fmt"

func matrixSumQueries(n int, queries [][]int) int64 {
	rowMap, colMap := make(map[int]bool), make(map[int]bool)
	row, col := n, n
	var res int64
	for i := len(queries) - 1; i >= 0; i-- {
		t, index, val := queries[i][0], queries[i][1], queries[i][2]
		if t == 0 && !rowMap[index] {
			rowMap[index] = true
			res += int64(row * val)
			col--
		} else if t == 1 && !colMap[index] {
			colMap[index] = true
			res += int64(col * val)
			row--
		}
	}
	return res
}

func main() {
	n, queries := 3, [][]int{{0,0,4}, {0,1,2}, {1,0,1}, {0,2,3}, {1,2,1}}
	res := matrixSumQueries(n, queries)
	fmt.Println(res)
}
