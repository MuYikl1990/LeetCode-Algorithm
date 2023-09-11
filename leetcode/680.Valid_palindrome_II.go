package main

import "fmt"

func validPalindromeII(s string) bool {
	var isValid func(string, int, int) bool
	isValid = func(s string, left, right int) bool {
		for left < right {
			if s[left] != s[right] {
				return false
			}
			left++
			right--
		}
		return true
	}

	left, right := 0, len(s) - 1
	for left < right {
		if s[left] != s[right] {
			return isValid(s, left + 1, right) || isValid(s, left, right - 1)
		}
		left++
		right--
	}
	return true
}

func main() {
	s := "abca"
	res := validPalindromeII(s)
	fmt.Println(res)
}
