package main

import "fmt"

// 圆环上有n个点，编号为0~n-1，从0出发，每次可以逆时针和顺时针走，走k步能回到0有多少种情况
func backToOrigin(n, k int) int {
	dp := make([][]int, k + 1)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	dp[0][0] = 1
	for i := 1; i <= k; i++ {
		for j := 0; j < n; j++ {
			dp[i][j] = dp[i - 1][(j - 1 + n) % n] + dp[i - 1][(j + 1) % n]
		}
	}
	return dp[k][0]
}

func main() {
	n, k := 13, 15
	res := backToOrigin(n, k)
	fmt.Println(res)
}
