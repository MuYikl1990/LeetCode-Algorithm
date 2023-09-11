package main

import "fmt"

func partitionLabels(s string) []int {
	char := make(map[byte]int)
	for i := range s {
		char[s[i]] = i
	}
	var res []int
	end, start := 0, 0
	for i := range s {
		end = max763(end, char[s[i]])
		if i == end {
			res = append(res, i - start + 1)
			start = i + 1
		}
	}
	return res
}

func max763(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	s := "ababcbacadefegdehijhklij"
	res := partitionLabels(s)
	fmt.Println(res)
}
