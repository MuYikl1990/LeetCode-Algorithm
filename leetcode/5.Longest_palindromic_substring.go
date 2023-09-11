package main

import "fmt"

func longestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}

	start, end := 0, 0
	for i := 0; i < len(s); i++ {
		start1, end1 := getPosition(s, i, i)
		if end1 - start1 > end - start {
			start = start1
			end = end1
		}
		start2, end2 := getPosition(s, i, i + 1)
		if end2 - start2 > end - start {
			start = start2
			end = end2
		}
	}
	return s[start:end+1]
}

func getPosition(s string, start, end int) (int, int) {
	for start >= 0 && end < len(s) && s[start] == s[end] {
		start--
		end++
	}
	return start+1, end-1
}

func main() {
	s := "babad"
	res := longestPalindrome(s)
	fmt.Println(res)
}
