package main

import "fmt"

func findDiagonalOrder(mat [][]int) []int {
	m, n := len(mat), len(mat[0])
	var res []int
	x, y := 0, 0
	for i := 0; i < m * n; i++ {
		res = append(res, mat[x][y])
		if (x + y) % 2 == 0 {
			if y == n - 1 {
				x++
			} else if x == 0 {
				y++
			} else {
				x--
				y++
			}
		} else {
			if x == m - 1 {
				y++
			} else if y == 0 {
				x++
			} else {
				x++
				y--
			}
		}
	}
	return res
}

func main() {
	mat := [][]int{{1,2,3}, {4,5,6}, {7,8,9}}
	res := findDiagonalOrder(mat)
	fmt.Println(res)
}
