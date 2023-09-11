package main

import (
	"fmt"
)

func numSmallerByFrequency(queries []string, words []string) []int {
	count := make([]int, 12)
	for _, word := range words {
		count[f(word)]++
	}
	for i := 10; i >= 0; i-- {
		count[i] += count[i + 1]
	}

	res := make([]int, len(queries))
	for i, query := range queries {
		res[i] = count[f(query) + 1]
	}
	return res
}

func f(s string) int {
	num, char := 0, 'z'
	for _, ch := range s {
		if ch < char {
			char = ch
			num = 1
		} else if ch == char {
			num++
		}
	}
	return num
}

func main() {
	queries, words := []string{"bbb", "cc"}, []string{"a", "aa", "aaa", "aaaa"}
	res := numSmallerByFrequency(queries, words)
	fmt.Println(res)
}
