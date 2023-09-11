package main

import (
	"fmt"
	"strconv"
)

func numDupDigitsAtMostN(n int) int {
	str := strconv.Itoa(n)
	length := len(str)
	dp := make([][]int, length)
	for i := range dp {
		dp[i] = make([]int, 1024)
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}

	var dfsMostN func(int, int, bool, bool) int
	dfsMostN = func(i int, mask int, isLimit bool, isNum bool) int {
		if i == length {
			if isNum {
				return 1
			}
			return 0
		}

		if !isLimit && dp[i][mask] != -1 {
			return dp[i][mask]
		}

		res := 0
		if !isNum {
			res = dfsMostN(i + 1, mask, false, false)
		}

		down := 0
		if !isNum {
			down = 1
		}

		up := 9
		if isLimit {
			up = int(str[i] - '0')
		}

		for num := down; num <= up; num++ {
			if mask >> num & 1 == 0 {
				res += dfsMostN(i + 1, mask | 1 << num, isLimit && num == up, true)
			}
		}

		if !isLimit {
			dp[i][mask] = res
		}
		return res
	}

	res := dfsMostN(0, 0, true, false)
	return n - res
}

func main() {
	n := 1000
	res := numDupDigitsAtMostN(n)
	fmt.Println(res)
}
