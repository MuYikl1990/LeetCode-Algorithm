package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"unicode"
)

func myAtoi(s string) int {
	s = strings.TrimSpace(s)
	if len(s) == 0 {
		return 0
	}
	if !unicode.IsDigit(rune(s[0])) && s[0] != '+' && s[0] != '-' {
		return 0
	}
	neg := s[0] == '-'
	i := 0
	res := 0
	if !unicode.IsDigit(rune(s[0])) {
		i = 1
	}
	for i < len(s) && unicode.IsDigit(rune(s[i])) {
		val, _ := strconv.Atoi(string(s[i]))
		res = res * 10 +val
		if !neg && res > math.MaxInt32 {
			return math.MaxInt32
		}
		if neg && -1 * res < math.MinInt32 {
			return math.MinInt32
		}
		i++
	}
	if neg {
		return -1 * res
	} else {
		return res
	}
}

func main() {
	s1 := "   -42"
	s2 := "4193 with words"
	s3 := " ++32 876"
	s4 := "+0023 "
	res1, res2, res3, res4 := myAtoi(s1), myAtoi(s2), myAtoi(s3), myAtoi(s4)
	fmt.Println(res1, res2, res3, res4)
}
