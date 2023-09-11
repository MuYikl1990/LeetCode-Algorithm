package main

import "fmt"

func kthSmallestInMatrix(matrix [][]int, k int) int {
	m, n := len(matrix), len(matrix[0])
	var checkCount func(int) bool
	checkCount = func(mid int) bool {
		i, j := m - 1, 0
		count := 0
		for i >= 0 && j < n {
			if matrix[i][j] <= mid {
				count += i + 1
				j++
			} else {
				i--
			}
		}
		return count >= k
	}

	left, right := matrix[0][0], matrix[m - 1][n - 1]
	for left < right {
		mid := left + (right - left) / 2
		if checkCount(mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

func main() {
	matrix, k := [][]int{{1,5,9},{10,11,13},{12,13,15}}, 6
	res := kthSmallestInMatrix(matrix, k)
	fmt.Println(res)
}
