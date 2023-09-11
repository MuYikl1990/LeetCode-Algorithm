package main

import (
	"fmt"
	"strings"
)

func validPalindrome(s string) bool {
	s = strings.ToLower(s)
	var isValid func(byte) bool
	isValid = func(char byte) bool {
		if (char >= 'a' && char <= 'z') || (char >= '0' && char <= '9') {
			return true
		}
		return false
	}

	left, right := 0, len(s) - 1
	for left < right {
		for left < right && !isValid(s[left]) {
			left++
		}
		for left < right && !isValid(s[right]) {
			right--
		}
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}

func main() {
	s := "A man, a plan, a canal: Panama"
	res := validPalindrome(s)
	fmt.Println(res)
}
