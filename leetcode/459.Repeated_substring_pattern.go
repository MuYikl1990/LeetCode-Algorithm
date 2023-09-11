package main

import (
	"fmt"
	"strings"
)

func repeatedSubstringPattern(s string) bool {
	return strings.Contains((s + s)[1:2 * len(s) - 1], s)
}

func main() {
	s := "abcabcabcabcv"
	res := repeatedSubstringPattern(s)
	fmt.Println(res)
}
