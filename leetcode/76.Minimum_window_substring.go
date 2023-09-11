package main

import (
	"fmt"
	"math"
)

func minWindow(s string, t string) string {
	sMap := make(map[byte]int)
	tMap := make(map[byte]int)
	sCount, tCount := 0, 0
	start, left, right, length := 0, 0, 0, math.MaxInt
	for i, _ := range t {
		if tMap[t[i]] == 0 {
			tCount++
		}
		tMap[t[i]]++
	}

	for right < len(s) {
		if _, ok := tMap[s[right]]; ok {
			sMap[s[right]]++
			if sMap[s[right]] == tMap[s[right]] {
				sCount++
			}
		}
		right++

		for sCount == tCount {
			if right - left < length {
				length = right - left
				start = left
			}

			if _, ok := tMap[s[left]]; ok {
				if sMap[s[left]] == tMap[s[left]] {
					sCount--
				}
				sMap[s[left]]--
			}
			left++
		}
	}

	if length == math.MaxInt {
		return ""
	}
	return s[start:start+length]
}

func main() {
	s, t := "ADOBECODEBANC", "ABC"
	res := minWindow(s, t)
	fmt.Println(res)
}
