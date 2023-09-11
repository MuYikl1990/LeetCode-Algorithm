package main

import "fmt"

// O(n^2)
func reverseParentheses(s string) string {
	var stack []int
	var res []byte
	sArr := []byte(s)

	for i := 0; i < len(sArr); i++ {
		if sArr[i] == '(' {
			stack = append(stack, i)
		} else if sArr[i] == ')' {
			start := stack[len(stack) - 1]
			reverse1190(sArr, start + 1, i - 1)
			stack = stack[:len(stack) - 1]
		}
	}

	for j := 0; j < len(sArr); j++ {
		if sArr[j] != '(' && sArr[j] != ')' {
			res = append(res, sArr[j])
		}
	}

	return string(res)
}

func reverse1190(sArr []byte, start, end int) {
	for start < end {
		sArr[start], sArr[end] = sArr[end], sArr[start]
		start++
		end--
	}
}

// O(n)
func reverseParentheses1(s string) string {
	var stack []int
	sArr := make([]int, len(s))
	var res []byte

	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stack = append(stack, i)
		} else if s[i] == ')' {
			index := stack[len(stack) - 1]
			sArr[i] = index
			sArr[index] = i
			stack = stack[:len(stack) - 1]
		}
	}

	for j, step := 0, 1; j < len(s); j += step {
		if s[j] == '(' || s[j] == ')' {
			j = sArr[j]
			step *= -1
		} else {
			res = append(res, s[j])
		}
	}
	return string(res)
}

func main() {
	s := "(ed(et(oc))el)"
	res := reverseParentheses(s)
	fmt.Println(res)
}
