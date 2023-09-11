package main

import "fmt"

func findMaxForm(strs []string, m int, n int) int {
	dp := make([][]int, m + 1)
	for i := range dp {
		dp[i] = make([]int, n + 1)
	}

	for _, str := range strs {
		num := zeroAndOne(str)
		for i := m; i >= num[0]; i-- {
			for j := n; j >= num[1]; j-- {
				dp[i][j] = max474(dp[i][j], dp[i - num[0]][j - num[1]] + 1)
			}
		}
	}
	return dp[m][n]
}

func max474(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func zeroAndOne(str string) [2]int {
	num := [2]int{}
	for i := range str {
		if str[i] == '0' {
			num[0]++
		} else {
			num[1]++
		}
	}
	return num
}

func main() {
	strs, m, n := []string{"10","0001","111001","1","0"}, 5, 3
	res := findMaxForm(strs, m, n)
	fmt.Println(res)
}
