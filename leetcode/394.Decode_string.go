package main

import (
	"fmt"
	"strconv"
)

func decodeString(s string) string {
	str, num := "", 0
	var strStack []string
	var numStack []int
	for i := range s {
		if s[i] == '[' {
			strStack = append(strStack, str)
			numStack = append(numStack, num)
			str, num = "", 0
		} else if s[i] == ']' {
			char := strStack[len(strStack) - 1]
			strStack = strStack[:len(strStack) - 1]
			number := numStack[len(numStack) - 1]
			numStack = numStack[:len(numStack) - 1]
			tmp := str
			for number > 1 {
				str += tmp
				number--
			}
			str = char + str
		} else if '0' <= s[i] && s[i] <= '9' {
			n, _ := strconv.Atoi(string(s[i]))
			num = num * 10 + n
		} else {
			str += string(s[i])
		}
	}
	return str
}

func main() {
	s := "3[a2[c]]"
	res := decodeString(s)
	fmt.Println(res)
}
