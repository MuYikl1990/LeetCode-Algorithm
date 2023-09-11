package main

import "fmt"

func lengthOfLastWord(s string) int {
	length := len(s)
	res, i := 0, length - 1
	for ; i >= 0; {
		if s[i] == ' ' {
			i--
		} else {
			for i >= 0 && s[i] != ' ' {
				res++
				i--
			}
			break
		}
	}
	return res
}

func main() {
	s := "   fly me   to   the moon  "
	res := lengthOfLastWord(s)
	fmt.Println(res)
}
