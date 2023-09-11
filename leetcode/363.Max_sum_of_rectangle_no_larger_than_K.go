package main

import (
	"fmt"
	"math"
)

func maxSumSubmatrix(matrix [][]int, k int) int {
	m, n := len(matrix), len(matrix[0])
	var dfs func([]int) int
	dfs = func(arr []int) int {
		max := math.MinInt
		for up := 0; up < m; up++ {
			sum := 0
			for down := up; down < m; down++ {
				sum += arr[down]
				if sum == k {
					return k
				}
				if sum > max && sum < k {
					max = sum
				}
			}
		}
		return max
	}

	res := math.MinInt
	for left := 0; left < n; left++ {
		arr := make([]int, m)
		for right := left; right < n; right++ {
			for i := 0; i < m; i++ {
				arr[i] += matrix[i][right]
			}
			res = max363(res, dfs(arr))
		}
	}
	return res
}

func max363(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	matrix, k := [][]int{{1,0,1}, {0,-2,3}}, 2
	res := maxSumSubmatrix(matrix, k)
	fmt.Println(res)
}
