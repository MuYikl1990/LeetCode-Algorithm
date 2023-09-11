package main

import (
	"fmt"
	"math"
)

func minFallingPathSum(grid [][]int) int {
	res := math.MaxInt
	s1, s2 := -1, -1
	dp := make([][]int, len(grid))
	for i := range dp {
		dp[i] = make([]int, len(grid[0]))
	}
	for i := 0; i < len(grid[0]); i++ {
		dp[0][i] = grid[0][i]
		if s1 == -1 {
			s1 = i
		} else {
			if dp[0][i] < dp[0][s1] {
				s2 = s1
				s1 = i
			} else if s2 == -1 || dp[0][i] < dp[0][s2] {
				s2 = i
			}
		}
	}

	for i := 1; i < len(grid); i++ {
		t1, t2 := -1, -1
		for j := 0; j < len(grid[0]); j++ {
			if j != s1 {
				dp[i][j] = dp[i - 1][s1] + grid[i][j]
			} else {
				dp[i][j] = dp[i - 1][s2] + grid[i][j]
			}

			if t1 == -1 {
				t1 = j
			} else {
				if dp[i][j] < dp[i][t1] {
					t2 = t1
					t1 = j
				} else if t2 == -1 || dp[i][j] < dp[i][t2] {
					t2 = j
				}
			}
		}
		s1, s2 = t1, t2
	}

	for i := 0; i < len(dp[0]); i++ {
		if dp[len(dp) - 1][i] < res {
			res = dp[len(dp) - 1][i]
		}
	}
	return res
}

func main() {
	grid := [][]int{{1,2,3}, {4,5,6}, {7,8,9}}
	res := minFallingPathSum(grid)
	fmt.Println(res)
}
