package main

import (
	"fmt"
	"strconv"
)

func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}

	n1, n2 := len(num1), len(num2)
	sum, carry := 0, 0
	res := ""

	for i := n1 - 1 + n2 - 1; i >= 0; i-- {
		sum = carry
		left, right := 0, i
		for left < n1 && right >= 0 {
			if left < n1 && right < n2 {
				sum += int(num1[left] - '0') * int(num2[right] - '0')
			}
			left++
			right--
		}
		res = strconv.Itoa(sum % 10) + res
		carry = sum / 10
	}
	if carry != 0 {
		res = strconv.Itoa(carry) + res
	}
	return res
}

func main() {
	num1, num2 := "123", "456"
	res := multiply(num1, num2)
	fmt.Println(res == "56088")
}
