package main

import "fmt"

func isValid(s string) bool {
	var stack []rune
	parentheses := map[rune]rune{')': '(', ']': '[', '}': '{'}
	for _, char := range s {
		if char == '(' || char == '[' || char == '{' {
			stack = append(stack, char)
		} else {
			if len(stack) == 0 {
				return false
			}
			cur := stack[len(stack)-1]
			if cur != parentheses[char] {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0

	//if len(s) % 2 == 1 {
	//	return false
	//}
	//stack := []byte{}
	//dickMap := map[byte]byte{')': '(', '}': '{', ']': '['}
	//for i := 0; i < len(s); i++ {
	//	if dickMap[s[i]] > 0 {
	//		if len(stack) == 0 || dickMap[s[i]] != stack[len(stack) - 1] {
	//			return false
	//		}
	//		stack = stack[:len(stack) - 1]
	//	} else {
	//		stack = append(stack, s[i])
	//	}
	//}
	//return len(stack) == 0
}

func main() {
	s := "()[]{(())}"
	res := isValid(s)
	fmt.Println(res)
}
