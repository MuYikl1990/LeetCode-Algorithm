package main

import (
	"fmt"
	"math"
	"strconv"
)


// 规律法
func countDigitOne(n int) int {
	res := 0
	for i := 1; i <= n; i *= 10 {
		res += n / (i * 10) * i + min233(max233(n % (i * 10) - i + 1, 0), i)
	}
	return res
}

func max233(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min233(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// 数位DP
func countDigitOneWithDp(n int) int {
	if n == 0 {
		return 0
	}
	stringN := strconv.Itoa(n)
	length := len(stringN)
	if length < 2 {
		return 1
	}
	dp1, dp2 := make([]int, length), make([]int, length)
	dp1[0] = 1
	for i := 1; i < length; i++ {
		dp1[i] = 10 * dp1[i - 1] + int((math.Pow10(i) * 10) / 10)
	}

	if stringN[length - 1] == '0' {
		dp2[0] = 0
	} else {
		dp2[0] = 1
	}

	for i, j := length - 2, 1; i >= 0; i, j = i - 1, j + 1 {
		cur := stringN[i]
		if cur == '0' {
			dp2[j] = dp2[j - 1]
		} else if cur == '1' {
			num, _ := strconv.Atoi(stringN[i + 1:])
			dp2[j] = dp1[j - 1] + dp2[j - 1] + num + 1
		} else {
			num, _ := strconv.Atoi(string(stringN[i]))
			dp2[j] = dp2[j - 1] + int((math.Pow10(j) * 10) / 10) + num * dp1[j - 1]
		}
	}
	return dp2[length - 1]
}

func main() {
	n := 1357
	res := countDigitOneWithDp(n)
	fmt.Println(res)
}
