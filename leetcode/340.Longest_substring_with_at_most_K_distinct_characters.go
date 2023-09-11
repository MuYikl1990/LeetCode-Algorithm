package main

import "fmt"

func LengthOfLongestSubstringKDistinct(s string, k int) int {
	if s == "" || k == 0 {
		return 0
	}
	charMap := make(map[byte]int)
	left, right := 0, 0
	res := 0
	for right < len(s) {
		cur := s[right]
		val := charMap[cur]
		if val == 0 {
			charMap[cur]++
			k--
		} else {
			charMap[cur]++
		}
		for k < 0 {
			charMap[s[left]]--
			if charMap[s[left]] == 0 {
				k++
			}
			left++
		}
		res = max340(res, right - left + 1)
		right++
	}
	return res
}

func max340(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	s, k := "ecebd", 3
	res := LengthOfLongestSubstringKDistinct(s, k)
	fmt.Println(res)
}
