package main

import "fmt"

func pathsWithMaxScore(board []string) []int {
	m, n := len(board), len(board[0])
	mod := 1000000007
	scoreDp, pathDp := make([][]int, m + 1), make([][]int, m + 1)
	for i := 0; i <= m; i++ {
		scoreDp[i] = make([]int, n + 1)
		pathDp[i] = make([]int, n + 1)
	}
	pathDp[m][n] = 1

	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if board[i][j] != 'X' && (pathDp[i + 1][j] != 0 || pathDp[i][j + 1] != 0 || pathDp[i + 1][j + 1] != 0) {
				maxScore := max1301(scoreDp[i + 1][j], max1301(scoreDp[i][j + 1], scoreDp[i + 1][j + 1]))
				scoreDp[i][j] = maxScore + int(board[i][j] - '0')
				if scoreDp[i + 1][j] == maxScore {
					pathDp[i][j] = (pathDp[i][j] + pathDp[i + 1][j]) % mod
				}
				if scoreDp[i][j + 1] == maxScore {
					pathDp[i][j] = (pathDp[i][j] + pathDp[i][j + 1]) % mod
				}
				if scoreDp[i + 1][j + 1] == maxScore {
					pathDp[i][j] = (pathDp[i][j] + pathDp[i + 1][j + 1]) % mod
				}
			}
		}
	}

	if scoreDp[0][0] == 0 {
		return []int{0, pathDp[0][0]}
	} else {
		return []int{scoreDp[0][0] - ('E' - '0') - ('S' - '0'), pathDp[0][0]}
	}
}

func max1301(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	board := []string{"E12","1X1","21S"}
	res := pathsWithMaxScore(board)
	fmt.Println(res)
}
