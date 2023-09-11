package main

import "fmt"

func removeInvalidParentheses(s string) []string {
	var res []string
	curMap := map[string]struct{}{s: {}}
	for {
		for str := range curMap {
			if isValidString(str) {
				res = append(res, str)
			}
		}
		if len(res) > 0 {
			return res
		}

		nextMap := make(map[string]struct{})
		for str := range curMap {
			for i := 0; i < len(str); i++ {
				if i > 0 && str[i] == str[i - 1] {
					continue
				}
				if str[i] == '(' || str[i] == ')' {
					nextMap[str[:i] + str[i + 1:]] = struct{}{}
				}
			}
		}
		curMap = nextMap
	}
}

func isValidString(str string) bool {
	count := 0
	for i := range str {
		if str[i] == '(' {
			count++
		} else if str[i] == ')' {
			count--
			if count < 0 {
				return false
			}
		}
	}
	return count == 0
}

func main() {
	s := "(a)())()"
	res := removeInvalidParentheses(s)
	fmt.Println(res)
}
