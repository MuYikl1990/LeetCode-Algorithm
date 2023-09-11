package main

import (
	"fmt"
)

func shiftingLetters(s string, shifts []int) string {
	cnt := 0
	runes := []rune(s)
	for i := len(s) - 1; i >= 0; i-- {
		cnt = (cnt + shifts[i]) % 26
		runes[i] = (runes[i] - 'a' + rune(cnt)) % 26 + 'a'
	}
	return string(runes)
}

func main() {
	s, shifts := "aaa", []int{1,2,3}
	res := shiftingLetters(s, shifts)
	fmt.Println(res)
}
