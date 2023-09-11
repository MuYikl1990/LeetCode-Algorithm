package main

import "fmt"

func countSubstrings(s string) int {
	res := 0
	for i := 0; i < len(s); i++ {
		for j := 0; j < 2; j++ {
			left, right := i, i + j
			for left >= 0 && right < len(s) && s[left] == s[right] {
				res++
				left--
				right++
			}
		}
	}
	return res
}

func main() {
	s := "ababa"
	res := countSubstrings(s)
	fmt.Println(res)
}
