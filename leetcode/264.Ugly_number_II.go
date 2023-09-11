package main

import "fmt"

func nthUglyNumber(n int) int {
	dp := make([]int, n + 1)
	dp[1] = 1
	p2, p3, p5 := 1, 1, 1
	for i := 2; i <= n; i++ {
		d2, d3, d5 := dp[p2] * 2, dp[p3] * 3, dp[p5] * 5
		dp[i] = min264(min264(d2, d3), d5)
		if dp[i] == d2 {
			p2++
		}
		if dp[i] == d3 {
			p3++
		}
		if dp[i] == d5 {
			p5++
		}
	}
	return dp[n]
}

func min264(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	n := 10
	res := nthUglyNumber(n)
	fmt.Println(res)
}
