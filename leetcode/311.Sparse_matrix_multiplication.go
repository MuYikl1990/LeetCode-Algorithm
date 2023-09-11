package main

import (
	"fmt"
)

func multiplySparseMatrix(a, b [][]int) [][]int {
	if len(a) == 0 || len(b) == 0 {
		return [][]int{}
	}

	m, n, l := len(a), len(a[0]), len(b[0])
	res := make([][]int, m)
	for p := range res {
		res[p] = make([]int, l)
	}

	for i := 0; i < m; i++ {
		for k := 0; k < n; k++ {
			if a[i][k] == 0 {
				continue
			}
			for j := 0; j < l; j++ {
				if b[k][j] == 0 {
					continue
				}
				res[i][j] = a[i][k] * b[k][j]
			}
		}
	}
	return res
}

func main() {
	a, b := [][]int{{1, 0, 0}, {-1, 0, 3}}, [][]int{{7, 0}, {0, 0}, {0, 1}}
	res := multiplySparseMatrix(a, b)
	fmt.Println(res)
}
