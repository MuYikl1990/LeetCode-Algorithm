package main

import (
	"fmt"
)

func solveSudoku(board [][]byte) {
	if len(board) == 0 {
		return
	}
	backtrace37(board)
}

func backtrace37(board [][]byte) bool {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] != '.' {
				continue
			}
			for k := 1; k < 10; k++ {
				if isValidNumber(board, i, j, k) {
					board[i][j] = byte('0' + k)
					if backtrace37(board) {
						return true
					}
					board[i][j] = '.'
				}
			}
			return false
		}
	}
	return true
}

func isValidNumber(board [][]byte, row, col, num int) bool {
	cell := byte('0' + num)
	for i := 0; i < len(board); i++ {
		if board[row][i] == cell {
			return false
		}
		if board[i][col] == cell {
			return false
		}
	}

	sBox, eBox := (row / 3) * 3, (col / 3) * 3
	for i := sBox; i < sBox + 3; i++ {
		for j := eBox; j < eBox + 3; j++ {
			if board[i][j] == cell {
				return false
			}
		}
	}
	return true
}

func main() {
	board := [][]byte{{'5','3','.','.','7','.','.','.','.'},
		{'6','.','.','1','9','5','.','.','.'},
		{'.','9','8','.','.','.','.','6','.'},
		{'8','.','.','.','6','.','.','.','3'},
		{'4','.','.','8','.','3','.','.','1'},
		{'7','.','.','.','2','.','.','.','6'},
		{'.','6','.','.','.','.','2','8','.'},
		{'.','.','.','4','1','9','.','.','5'},
		{'.','.','.','.','8','.','.','7','9'}}
	solveSudoku(board)
	fmt.Println(board)
}
