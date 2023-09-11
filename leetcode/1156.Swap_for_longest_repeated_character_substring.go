package main

import "fmt"

func maxRepOpt1(text string) int {
	word := make(map[uint8]int)
	for i := range text {
		word[text[i]]++
	}
	res, count, cur := 1, 1, text[0]
	for i := 1; i < len(text); i++ {
		if text[i] != cur {
			index := i + 1
			for index < len(text) && text[index] == cur {
				index++
				count++
			}

			if count < word[cur] {
				count++
			}
			res = max1156(res, count)
			cur = text[i]
			count = 1
		} else {
			count++
		}
	}

	if count > 1 {
		if count < word[cur] {
			count++
		}
		res = max1156(res, count)
	}
	return res
}

func max1156(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	text := "ababa"
	res := maxRepOpt1(text)
	fmt.Println(res)
}
