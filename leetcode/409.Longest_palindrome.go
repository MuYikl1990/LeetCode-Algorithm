package main

import "fmt"

func longestPalindromeI(s string) int {
	sMap := make(map[byte]int)
	res := 0
	for i := range s {
		sMap[s[i]]++
		if sMap[s[i]] == 2 {
			res += 2
			sMap[s[i]] = 0
		}
	}
	if res < len(s) {
		res++
	}
	return res
}

func main() {
	s := "abccccdd"
	res := longestPalindromeI(s)
	fmt.Println(res)
}
