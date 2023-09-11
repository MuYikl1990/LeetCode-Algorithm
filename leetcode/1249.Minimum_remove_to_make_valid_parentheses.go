package main

import "fmt"

func minRemoveToMakeValid(s string) string {
	var stack []int
	var res []byte
	invalid := make([]bool, len(s))

	for i := range s {
		if s[i] == '(' {
			stack = append(stack, i)
			invalid[i] = true
		} else if s[i] == ')' {
			if len(stack) == 0 {
				invalid[i] = true
			} else {
				invalid[stack[len(stack) - 1]] = false
				stack = stack[:len(stack) - 1]
			}
		}
	}

	for i := range invalid {
		if !invalid[i] {
			res = append(res, s[i])
		}
	}
	return string(res)
}

func main() {
	s := "lee(t(c)o)de)"
	res := minRemoveToMakeValid(s)
	fmt.Println(res)
}
