package main

import (
	"fmt"
)

func reorganizeString(s string) string {
	var chars [26]int
	list := []rune(s)
	max, half := 0, (len(list) + 1) / 2
	var char rune
	for i := range list {
		index := list[i] - 'a'
		chars[index]++
		if chars[index] > max {
			max = chars[index]
			char = list[i]
		}
		if max > half {
			return ""
		}
	}

	tmp, idx := make([]rune, len(list)), 0
	for chars[char - 'a'] > 0 {
		tmp[idx] = char
		chars[char - 'a']--
		idx += 2
	}

	for i := 0; i < len(list); i++ {
		for chars[list[i] - 'a'] > 0 {
			if idx >= len(list) {
				idx = 1
			}
			tmp[idx] = list[i]
			chars[list[i] - 'a']--
			idx += 2
		}
	}
	return string(tmp)
}

func main() {
	s := "bfrbs"
	res := reorganizeString(s)
	fmt.Println(res)
}
