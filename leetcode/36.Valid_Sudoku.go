package main

import "fmt"

func isValidSudoku(board [][]byte) bool {
	var row, col [9][9]int
	var box [3][3][9]int

	for i, r := range board {
		for j, c := range r {
			if c == '.' {
				continue
			}
			num := c - '1'
			row[i][num]++
			col[j][num]++
			box[i/3][j/3][num]++
			if row[i][num] > 1 || col[j][num] > 1 || box[i/3][j/3][num] > 1 {
				return false
			}
		}
	}
	return true
}

func main() {
	board := [][]byte{{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'8', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'}}
	res := isValidSudoku(board)
	fmt.Println(res)
}
