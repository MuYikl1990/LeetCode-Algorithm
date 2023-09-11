package main

import (
	"fmt"
	"strconv"
)

func addStrings(num1 string, num2 string) string {
	i, j, res := len(num1) - 1, len(num2) - 1, ""
	var tmp, carry uint8
	for i >= 0 || j >= 0 {
		tmp = 0
		if i >= 0 {
			tmp += num1[i] - '0'
		}
		if j >= 0 {
			tmp += num2[j] - '0'
		}
		tmp += carry
		carry = tmp / 10
		res = strconv.Itoa(int(tmp % 10)) + res
		i--
		j--
	}
	if carry != 0 {
		return "1" + res
	} else {
		return res
	}
}

func main() {
	num1, num2 := "31", "123"
	res := addStrings(num1, num2)
	fmt.Println(res)
}
