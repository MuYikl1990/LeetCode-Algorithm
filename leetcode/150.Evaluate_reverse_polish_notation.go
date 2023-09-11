package main

import (
	"fmt"
	"strconv"
)

func evalRPN(tokens []string) int {
	var stack []int
	for _, token := range tokens {
		if num, err := strconv.Atoi(token); err == nil {
			stack = append(stack, num)
		} else {
			cur := stack[len(stack) - 1]
			pre := stack[len(stack) - 2]
			stack = stack[:len(stack) - 2]
			if token == "+" {
				stack = append(stack, pre + cur)
			}
			if token == "-" {
				stack = append(stack, pre - cur)
			}
			if token == "*" {
				stack = append(stack, pre * cur)
			}
			if token == "/" {
				stack = append(stack, pre / cur)
			}
		}
	}
	return stack[0]
}

func main() {
	tokens := []string{"10","6","9","3","+","-11","*","/","*","17","+","5","+"}
	res := evalRPN(tokens)
	fmt.Println(res)
}
