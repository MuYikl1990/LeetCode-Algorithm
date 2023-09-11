package main

import "fmt"

func countRoutes(locations []int, start int, finish int, fuel int) int {
	if abs1575(locations[start], locations[finish]) > fuel {
		return 0
	}

	mod := 1000000007
	dp := make([][]int, len(locations))
	for i := range dp {
		dp[i] = make([]int, fuel + 1)
	}
	for i := 0; i <= fuel; i++ {
		dp[finish][i] = 1
	}

	for i := 0; i <= fuel; i++ {
		for j := 0; j < len(locations); j++ {
			for k := 0; k < len(locations); k++ {
				if j != k {
					need := abs1575(locations[j], locations[k])
					if i >= need {
						dp[j][i] = (dp[j][i] + dp[k][i - need]) % mod
					}
				}
			}
		}
	}
	return dp[start][fuel]
}

func abs1575(a, b int) int {
	if a < b {
		return b - a
	}
	return a - b
}

func main() {
	locations, start, finish, fuel := []int{4,3,1}, 1, 0, 6
	res := countRoutes(locations, start, finish, fuel)
	fmt.Println(res)
}
