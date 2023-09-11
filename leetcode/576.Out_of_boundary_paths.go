package main

import "fmt"

func findPaths(m int, n int, maxMove int, startRow int, startColumn int) int {
	mod, dirs := 1000000007, [5]int{-1, 0, 1, 0, -1}
	dp := make([][][]int, m)
	for i := range dp {
		dp[i] = make([][]int, n)
		for j := range dp[i] {
			dp[i][j] = make([]int, maxMove + 1)
		}
	}

	for k := 1; k <= maxMove; k++ {
		for i := 0; i < m; i++ {
			for j := 0; j < n; j++ {
				if i == 0 {
					dp[i][j][k]++
				}
				if i == m - 1 {
					dp[i][j][k]++
				}
				if j == 0 {
					dp[i][j][k]++
				}
				if j == n - 1 {
					dp[i][j][k]++
				}

				for d := 0; d < 4; d++ {
					ni := i + dirs[d]
					nj := j + dirs[d + 1]
					if ni >= 0 && nj >= 0 && ni < m && nj < n {
						dp[i][j][k] = (dp[i][j][k] + dp[ni][nj][k - 1]) % mod
					}
				}
			}
		}
	}
	return dp[startRow][startColumn][maxMove]
}

func main() {
	m, n, maxMove, startRow, startColumn := 1, 3, 3, 0, 1
	res := findPaths(m, n, maxMove, startRow, startColumn)
	fmt.Println(res)
}
