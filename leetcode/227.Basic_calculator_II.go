package main

import (
	"fmt"
	"strconv"
	"unicode"
)

func calculateII(s string) int {
	var stack []int
	op, num := '+', 0

	for i, char := range s {
		if unicode.IsDigit(char) {
			digit, _ := strconv.Atoi(string(char))
			num = 10 * num + digit
		}
		if i == len(s) - 1 || char == '+' || char == '-' || char == '*' || char == '/' {
			if op == '+' {
				stack = append(stack, num)
			} else if op == '-' {
				stack = append(stack, -num)
			} else if op == '*' {
				mul := stack[len(stack) - 1]
				stack = stack[:len(stack) - 1]
				stack = append(stack, mul * num)
			} else if op == '/' {
				div := stack[len(stack) - 1]
				stack = stack[:len(stack) - 1]
				stack = append(stack, div / num)
			}
			op = char
			num = 0
		}
	}

	res := 0
	for i := range stack {
		res += stack[i]
	}
	return res
}

func main() {
	s := "3+2*2"
	res := calculateII(s)
	fmt.Println(res)
}
