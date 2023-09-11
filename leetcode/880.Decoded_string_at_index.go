package main

import (
	"fmt"
	"unicode"
)

func decodeAtIndex(s string, k int) string {
	n := 0

	for _, str := range s {
		if unicode.IsDigit(str) {
			n *= int(str - '0')
		} else {
			n++
		}
	}

	for i := len(s) - 1; i >= 0; i-- {
		k %= n
		if k == 0 && !unicode.IsDigit(rune(s[i])) {
			return string(s[i])
		}
		if unicode.IsDigit(rune(s[i])) {
			n /= int(s[i] - '0')
		} else {
			n--
		}
	}
	return ""
}

func main() {
	s, k := "y959q969u3hb22odq595", 222280369
	res := decodeAtIndex(s, k)
	fmt.Println(res)
}
