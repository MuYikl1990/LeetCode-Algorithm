package main

import "fmt"

func totalNQueens(n int) int {
	res := 0
	var dfs52 func(int, int, int, int, int)
	dfs52 = func(n, row, col, ld, rd int) {
		if row >= n {
			res++
			return
		}
		bits := ^(col | ld | rd) & ((1 << n) - 1)
		for bits > 0 {
			pick := bits & -bits
			dfs52(n, row + 1, col | pick, (ld | pick) << 1, (rd | pick) >> 1)
			bits &= bits - 1
		}
	}
	dfs52(n, 0, 0, 0, 0)
	return res
}

func main() {
	n := 8
	res := totalNQueens(n)
	fmt.Println(res)
}
