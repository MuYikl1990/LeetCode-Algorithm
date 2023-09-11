package main

import (
	"fmt"
	"strconv"
)

func addBinary(a string, b string) string {
	res := ""
	carry := 0
	for i, j := len(a) - 1, len(b) - 1; i >= 0 || j >= 0 || carry > 0; i, j = i - 1, j - 1 {
		x, y := 0, 0
		if i >= 0 {
			x = int(a[i] - '0')
		}
		if j >= 0 {
			y = int(b[j] - '0')
		}
		sum := x + y + carry
		res = strconv.Itoa(sum % 2) + res
		carry = sum / 2
	}
	return res
}

func main() {
	a, b := "1010", "1011"
	res := addBinary(a, b)
	fmt.Println(res)
}
