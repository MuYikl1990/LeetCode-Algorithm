package main

import (
	"fmt"
	"sort"
)

func longestStrChain(words []string) int {
	sort.Slice(words, func(i, j int) bool {
		return len(words[i]) < len(words[j])
	})
	res := 1
	stringMap := make(map[string]int)
	for _, word := range words {
		if _, has := stringMap[word]; !has {
			stringMap[word] = 1
		}
		for i := range word {
			str := word[:i] + word[i + 1:]
			if _, ok := stringMap[str]; ok {
				stringMap[word] = max1048(stringMap[word], stringMap[str] + 1)
			}
		}
		res = max1048(res, stringMap[word])
	}
	return res
}

func max1048(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	words := []string{"xbc","pcxbcf","xb","cxbc","pcxbc"}
	res := longestStrChain(words)
	fmt.Println(res)
}
