package main

import "fmt"

func removeDuplicatesInString(s string) string {
	if len(s) < 2 {
		return s
	}
	var stack []byte
	for i := 0; i < len(s); i++ {
		if len(stack) > 0 && stack[len(stack) - 1] == s[i] {
			stack = stack[:len(stack) - 1]
		} else {
			stack = append(stack, s[i])
		}
	}
	return string(stack)
}

func main() {
	s := "aababaab"
	res := removeDuplicatesInString(s)
	fmt.Println(res)
}
