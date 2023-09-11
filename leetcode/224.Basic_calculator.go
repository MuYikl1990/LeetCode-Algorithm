package main

import (
	"fmt"
	"strconv"
	"unicode"
)

func calculate(s string) int {
	res, sign, num := 0, 1, 0
	var stack []int
	for _, char := range s {
		if char == ' ' {
			continue
		} else if unicode.IsDigit(char) {
			digit, _ := strconv.Atoi(string(char))
			num = 10 * num + digit
		} else if char == '+' || char == '-' {
			res += num * sign
			num = 0
			if char == '+' {
				sign = 1
			} else {
				sign = -1
			}
		} else if char == '(' {
			stack = append(stack, res)
			stack = append(stack, sign)
			res = 0
			sign = 1
		} else if char == ')' {
			res += sign * num
			num = 0
			res *= stack[len(stack)-1]
			res += stack[len(stack)-2]
			stack = stack[:len(stack)-2]
		}
	}
	res += sign * num
	return res
}

func main() {
	s := "(1-(4+5-2)- 3)+(6+8)"
	res := calculate(s)
	fmt.Println(res)
}
