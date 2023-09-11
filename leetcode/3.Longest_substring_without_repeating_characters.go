package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	if s == "" {
		return 0
	}

	sMap := make(map[string]int)
	maxLength, start := 0, 0
	for i := 0; i < len(s); i++ {
		str := string(s[i])
		if pos, ok := sMap[str]; ok {
			start = Max(start, pos)
		}
		maxLength = Max(maxLength, i - start + 1)
		sMap[str] = i + 1
	}
	return maxLength
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	s := " "
	res := lengthOfLongestSubstring(s)
	fmt.Println(res)
}
