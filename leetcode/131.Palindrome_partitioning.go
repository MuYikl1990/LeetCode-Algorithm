package main

import "fmt"

func partitionPalindrome(s string) [][]string {
	n := len(s)
	var res [][]string
	dp := make([][]bool, n)
	for i := range dp {
		dp[i] = make([]bool, n)
	}

	for j := 0; j < n; j++ {
		for i := 0; i <= j; i++ {
			if i == j {
				dp[i][j] = true
			} else if j - i == 1 && s[i] == s[j] {
				dp[i][j] = true
			} else if j - i > 1 && s[i] == s[j] && dp[i + 1][j - 1] {
				dp[i][j] = true
			}
		}
	}

	var backtrace func(string, int, []string)
	backtrace = func(s string, start int, strs []string) {
		if start == len(s) {
			tmp := make([]string, len(strs))
			copy(tmp, strs)
			res = append(res, tmp)
			return
		}

		for i := start; i < len(s); i++ {
			if dp[start][i] {
				strs = append(strs, s[start:i + 1])
				backtrace(s, i + 1, strs)
				strs = strs[:len(strs) - 1]
			}
		}
		return
	}
	backtrace(s, 0, []string{})
	return res
}

func main() {
	s := "abaabab"
	res := partitionPalindrome(s)
	fmt.Println(res)
}
