package main

import "fmt"

func longestValidParentheses(s string) int {
	left, right, res := 0, 0, 0
	for _, char := range s {
		if char == '(' {
			left++
		}
		if char == ')' {
			right++
		}
		if left == right {
			res = max32(res, 2 * right)
		} else if left < right {
			left, right = 0, 0
		}
	}
	left, right = 0, 0
	for i := len(s)-1; i > 0; i-- {
		if s[i] == '(' {
			left++
		}
		if s[i] == ')' {
			right++
		}
		if left == right {
			res = max32(res, 2 * left)
		} else if left > right {
			left, right = 0, 0
		}
	}
	return res
}

func max32(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	s := ")()())"
	res := longestValidParentheses(s)
	fmt.Println(res)
}
