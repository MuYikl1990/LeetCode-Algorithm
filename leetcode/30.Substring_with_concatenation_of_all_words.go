package main

import "fmt"

func findSubstring(s string, words []string) []int {
	n, wordNum, wordLen := len(s), len(words), len(words[0])
	wordMap := make(map[string]int)
	var res []int
	for _, word := range words {
		wordMap[word]++
	}

	for i := 0; i < wordLen; i++ {
		count, tmpMap := 0, make(map[string]int)
		left, right := i, i
		for right <= n - wordLen {
			word := s[right:right + wordLen]
			right += wordLen
			if _, exist := wordMap[word]; !exist {
				count = 0
				left = right
				tmpMap = make(map[string]int)
			} else {
				tmpMap[word]++
				count++
				for tmpMap[word] > wordMap[word] {
					removeWord := s[left:left + wordLen]
					count--
					tmpMap[removeWord]--
					left += wordLen
				}
				if count == wordNum {
					res = append(res, left)
				}
			}
		}
	}
	return res
}

func main() {
	s, words := "baaacbabba", []string{"ba","bb"}
	res := findSubstring(s, words)
	fmt.Println(res)
}
