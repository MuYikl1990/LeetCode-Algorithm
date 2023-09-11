package main

import "fmt"

func generateMatrix(n int) [][]int {
	res := make([][]int, n)
	for i := range res {
		res[i] = make([]int, n)
	}
	up, bottom, left, right := 0, n - 1, 0, n - 1
	num := 1
	for num <= n * n {
		for i := left; i <= right; i++ {
			res[up][i] = num
			num++
		}
		up++
		for i := up; i <= bottom; i++ {
			res[i][right] = num
			num++
		}
		right--
		for i := right; i >= left; i-- {
			res[bottom][i] = num
			num++
		}
		bottom--
		for i := bottom; i >= up; i-- {
			res[i][left] = num
			num++
		}
		left++
	}
	return res
}

func main() {
	n := 3
	res := generateMatrix(n)
	fmt.Println(res)
}
