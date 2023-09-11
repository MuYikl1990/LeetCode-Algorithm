package main

import (
	"fmt"
	"strconv"
)

func minusStrings(string1, string2 string) string {
	if len(string1) > len(string2) || (len(string1) == len(string2) && string1 > string2) {
		return subString(string1, string2)
	}
	return "-" + subString(string2, string1)
}

func subString(a, b string) string {
	res := ""
	i, j := len(a) - 1, len(b) - 1
	var borrow uint8
	for i >= 0 || j >= 0 {
		var x, y uint8
		if i >= 0 {
			x = a[i] - '0'
		}
		if j >= 0 {
			y = b[j] - '0'
		}
		ans := (x - borrow - y + 10) % 10
		res = strconv.Itoa(int(ans)) + res
		if x - borrow - y + 10 < 10 {
			borrow = 1
		} else {
			borrow = 0
		}
		i--
		j--
	}
	k := 0
	for ;k < len(res) - 1; k++ {
		if res[k] != '0' {
			break
		}
	}
	return res[k:]
}

func main() {
	string1, string2 := "110", "321"
	res := minusStrings(string1, string2)
	fmt.Println(res)
}
