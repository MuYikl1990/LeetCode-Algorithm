package main

import "fmt"

func findAnagrams(s string, p string) []int {
	sList, pList := [26]int{}, [26]int{}
	for _, char := range p {
		pList[char - 'a']++
	}

	left, right := 0, 0
	m, n := len(s), len(p)
	var res []int
	for right < m {
		cur := s[right]
		sList[cur - 'a']++
		right++
		for sList[cur - 'a'] > pList[cur - 'a'] {
			sList[s[left] - 'a']--
			left++
		}
		if right - left == n {
			res = append(res, left)
		}
	}
	return res
}

func main() {
	s, p := "cbaebabacd", "abc"
	res := findAnagrams(s, p)
	fmt.Println(res)
}
