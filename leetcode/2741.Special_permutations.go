package main

import "fmt"

// 状态压缩dp
func specialPerm(nums []int) int {
	mod := int(1e9 + 7)
	n := len(nums)
	m := 1 << n
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	for i := range dp[0] {
		dp[0][i] = 1
	}

	for i := 1; i < m; i++ {
		for j, prev := range nums {
			for k, cur := range nums {
				if (i >> j) & 1 == 1 && j != k && (prev % cur == 0 || cur % prev == 0) {
					dp[i][k] = (dp[i][k] + dp[i ^ (1 << j)][j]) % mod
				}
			}
		}
	}

	res := 0
	for i := 0; i < n; i++ {
		res = (res + dp[(m - 1) ^ (1 << i)][i]) % mod
	}
	return res
}

func main() {
	nums := []int{20,100,50,5,10,70,7}
	res := specialPerm(nums)
	fmt.Println(res)
}
