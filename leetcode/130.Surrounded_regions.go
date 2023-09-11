package main

import "fmt"

func solve(board [][]byte)  {
	if len(board) == 0 {
		return
	}

	m, n := len(board), len(board[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if (i == 0 || i == m - 1 || j == 0 || j == n - 1) && board[i][j] == 'O' {
				dfs130(board, i, j, m, n)
			}
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 'O' {
				board[i][j] = 'X'
			}
			if board[i][j] == '#' {
				board[i][j] = 'O'
			}
		}
	}
	return
}

func dfs130(board [][]byte, i, j, m, n int) {
	if i < 0 || j < 0 || i >= m || j >= n || board[i][j] == 'X' || board[i][j] == '#' {
		return
	}

	board[i][j] = '#'
	dir := [5]int{-1, 0, 1, 0, -1}
	for k := 0; k < 4; k++ {
		x := i + dir[k]
		y := j + dir[k + 1]
		dfs130(board, x, y, m, n)
	}
}

func main() {
	board := [][]byte{{'X','X','X','X'}, {'X','O','O','X'}, {'X','X','O','X'}, {'X','O','X','X'}}
	solve(board)
	fmt.Println(board)
}
