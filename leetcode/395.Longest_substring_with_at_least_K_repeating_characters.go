package main

import "fmt"

func longestSubstring(s string, k int) int {
	var dfs func(int, int) int
	dfs = func(start, end int) int {
		if end - start + 1 < k {
			return 0
		}
		charMap := make(map[byte]int)
		for i := start; i <= end; i++ {
			charMap[s[i]]++
		}

		for end - start + 1 >= k && charMap[s[start]] < k {
			start++
		}

		for end - start + 1 >= k && charMap[s[end]] < k {
			end--
		}

		if end - start + 1 < k {
			return 0
		}

		for i := start; i <= end; i++ {
			if charMap[s[i]] < k {
				return max395(dfs(start, i - 1), dfs(i + 1, end))
			}
		}
		return end - start + 1
	}
	return dfs(0, len(s) - 1)
}

func max395(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	s, k := "ababbc", 2
	res := longestSubstring(s, k)
	fmt.Println(res)
}
