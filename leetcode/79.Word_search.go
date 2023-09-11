package main

import "fmt"

func exist(board [][]byte, word string) bool {
	m, n := len(board), len(board[0])
	used := make([][]bool, m)
	for i := range used {
		used[i] = make([]bool, n)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == word[0] && canFind(i, j, 0, used, word, board) {
				return true
			}
		}
	}
	return false
}

func canFind(row, col, index int, used [][]bool, word string, board [][]byte) bool {
	if index == len(word) {
		return true
	}

	if row < 0 || row >= len(used) || col < 0 || col >= len(used[0]) {
		return false
	}

	if used[row][col] || board[row][col] != word[index] {
		return false
	}

	used[row][col] = true
	flag := canFind(row - 1, col, index + 1, used, word, board) || canFind(row + 1, col, index + 1, used, word, board) ||
		canFind(row, col - 1, index + 1, used, word, board) || canFind(row, col + 1, index + 1, used, word, board)

	if flag {
		return true
	} else {
		used[row][col] = false
		return false
	}
}

func main() {
	board, word := [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'X', 'E', 'E'}}, "ABCCED"
	res := exist(board, word)
	fmt.Println(res)
}
