package main

import (
	"fmt"
)

// 数位dp
func count(num1 string, num2 string, minSum int, maxSum int) int {
	mod := 1000000007
	countIntegers := func(str string) int {
		length := len(str)
		dp := make([][]int, length)
		for i := range dp {
			dp[i] = make([]int, maxSum + 1)
			for j := range dp[i] {
				dp[i][j] = -1
			}
		}

		var dfs func(int, int, bool) int
		dfs = func(i int, sum int, isLimit bool) int {
			if sum > maxSum {
				return 0
			}

			if i == length {
				if sum >= minSum {
					return 1
				}
				return 0
			}

			if !isLimit && dp[i][sum] != -1 {
				return dp[i][sum]
			}

			up := 9
			if isLimit {
				up = int(str[i] - '0')
			}

			res := 0
			for d := 0; d <= up; d++ {
				res = (res + dfs(i + 1, sum + d, isLimit && d == up)) % mod
			}

			if !isLimit {
				dp[i][sum] = res
			}
			return res
		}

		return dfs(0, 0, true)
	}

	if num1 == "1" {
		return countIntegers(num2)
	}

	res := (countIntegers(num2) - countIntegers(num1) + mod) % mod
	sum := 0
	for i := range num1 {
		sum += int(num1[i] - '0')
	}
	if minSum <= sum && sum <= maxSum {
		res++
	}
	return res
}

func min2719(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	num1, num2, minSum, maxSum := "1012640017461217236611", "9234552128261772272769", 67, 148
	res := count(num1, num2, minSum, maxSum)
	fmt.Println(res)
}
