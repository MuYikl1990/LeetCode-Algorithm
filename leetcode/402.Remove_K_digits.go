package main

import (
	"fmt"
	"strings"
)

func removeKdigits(num string, k int) string {
	var stack []string
	remain := len(num) - k

	for _, char := range num {
		for k > 0 && len(stack) > 0 && stack[len(stack) - 1] > string(char) {
			stack = stack[:len(stack) - 1]
			k--
		}
		stack = append(stack, string(char))
	}

	stack = stack[:remain]
	res := strings.Join(stack, "")
	res = strings.TrimLeft(res, "0")
	if res == "" {
		return "0"
	}
	return res
}

func main() {
	num, k := "1000", 1
	res := removeKdigits(num, k)
	fmt.Println(res)
}
