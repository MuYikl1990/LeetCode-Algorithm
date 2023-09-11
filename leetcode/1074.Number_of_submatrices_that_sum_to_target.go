package main

import "fmt"

func numSubmatrixSumTarget(matrix [][]int, target int) int {
	m, n := len(matrix), len(matrix[0])
	sum := make([][]int, m + 1)
	for i := range sum {
		sum[i] = make([]int, n + 1)
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			sum[i][j] = sum[i - 1][j] + sum[i][j - 1] - sum[i - 1][j - 1] + matrix[i - 1][j - 1]
		}
	}

	res := 0
	for i := 1; i <= m; i++ {
		for j := i; j <= m; j++ {
			sumMap := make(map[int]int)
			for col := 1; col <= n; col++ {
				cur := sum[j][col] - sum[i - 1][col]
				if cur == target {
					res++
				}
				if val, ok := sumMap[cur - target]; ok {
					res += val
				}
				sumMap[cur]++
			}
		}
	}
	return res
}

func main() {
	matrix, target := [][]int{{0,1,0}, {1,1,1}, {0,1,0}}, 0
	res := numSubmatrixSumTarget(matrix, target)
	fmt.Println(res)
}
