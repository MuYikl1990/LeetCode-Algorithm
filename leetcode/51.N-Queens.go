package main

import (
	"fmt"
	"strings"
)

func solveNQueens(n int) [][]string {
	if n == 0 {
		return [][]string{}
	}

	board := make([][]string, n)
	for i := 0; i < n; i++ {
		board[i] = make([]string, n)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			board[i][j] = "."
		}
	}
	var res [][]string
	backtrace51(board, 0, n, &res)
	return res
}

func backtrace51(board [][]string, row, n int, res *[][]string) {
	if row == n {
		var temp []string
		for i := range board {
			str := strings.Join(board[i], "")
			temp = append(temp, str)
		}
		*res = append(*res, temp)
		return
	}

	for col := 0; col < n; col++ {
		if !isValid51(board, row, col) {
			continue
		}
		board[row][col] = "Q"
		backtrace51(board, row + 1, n, res)
		board[row][col] = "."
	}
	return
}

func isValid51(board [][]string, row, col int) bool {
	for i := 0; i < len(board); i++ {
		if board[i][col] == "Q" {
			return false
		}
	}

	i := row - 1
	j := col - 1
	for i >= 0 && j >= 0 {
		if board[i][j] == "Q" {
			return false
		}
		i--
		j--
	}

	i = row - 1
	j = col + 1
	for i >= 0 && j < len(board) {
		if board[i][j] == "Q" {
			return false
		}
		i--
		j++
	}
	return true
}

func main() {
	n := 6
	res := solveNQueens(n)
	fmt.Println(res)
}
