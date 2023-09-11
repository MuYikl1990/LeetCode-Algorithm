package main

import (
	"fmt"
	"math"
)

func getMaxMatrix(matrix [][]int) []int {
	m, n := len(matrix), len(matrix[0])
	start, dp, max := 0, 0, math.MinInt
	sum, res := make([]int, n), make([]int, 4)

	for i := 0; i < m; i++ {
		sum = make([]int, n)
		for j := i; j < m; j++ {
			start = 0
			dp = 0
			for k := 0; k < n; k++ {
				sum[k] += matrix[j][k]
				dp += sum[k]

				if dp > max {
					max = dp
					res[0] = i
					res[1] = start
					res[2] = j
					res[3] = k
				}

				if dp < 0 {
					dp = 0
					start = k + 1
				}
			}
		}
	}
	return res
}

func main() {
	matrix := [][]int{{1,2,3,5}, {-1,2,-4,-3}, {0,-1,7,-1}}
	res := getMaxMatrix(matrix)
	fmt.Println(res)
}
