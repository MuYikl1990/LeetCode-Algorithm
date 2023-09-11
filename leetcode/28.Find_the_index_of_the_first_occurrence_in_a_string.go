package main

import "fmt"

// KMP
func strStr(haystack string, needle string) int {
	next := buildNext(needle)
	i, j := 0, 0
	m, n := len(haystack), len(needle)
	for i < m && j < n {
		if j == -1 || haystack[i] == needle[j] {
			i++
			j++
		} else {
			j = next[j]
		}
	}
	if j == n {
		return i - j
	}
	return -1
}

func buildNext(needle string) []int {
	n := len(needle)
	next := make([]int, n + 1)
	next[0] = -1
	i, j := 0, -1
	for i < n {
		if j == -1 || needle[i] == needle[j] {
			i++
			j++
			next[i] = j
		} else {
			j = next[j]
		}
	}
	return next
}

func main() {
	haystack, needle := "aleetcode", "leetco"
	res := strStr(haystack, needle)
	fmt.Println(res)
}
