package main

import (
	"fmt"
	"strings"
)

func wordBreakII(s string, wordDict []string) []string {
	wordMap := make(map[string]struct{})
	for _, word := range wordDict {
		wordMap[word] = struct{}{}
	}

	dp := make([][][]string, len(s))
	var res []string
	var dfs func(int) [][]string
	dfs = func(index int) [][]string {
		if index >= len(s) {
			return [][]string{}
		}
		if dp[index] != nil {
			return dp[index]
		}

		var list [][]string
		for i := index + 1; i < len(s); i++ {
			word := s[index:i]
			if _, has := wordMap[word]; has {
				for _, ans := range dfs(i) {
					ans = append([]string{word}, ans...)
					list = append(list, ans)
				}
			}
		}

		word := s[index:]
		if _, has := wordMap[word]; has {
			list = append(list, []string{word})
		}

		dp[index] = list
		return list
	}

	wordList := dfs(0)
	for i := range wordList {
		res = append(res, strings.Join(wordList[i], " "))
	}
	return res
}

func main() {
	s, wordDict := "catsanddog", []string{"cats","dog","sand","and","cat"}
	res := wordBreakII(s, wordDict)
	fmt.Println(res)  // ["cats and dog","cat sand dog"]
}
