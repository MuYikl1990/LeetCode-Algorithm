package main

import "fmt"

func checkInclusion(s1 string, s2 string) bool {
	m, n := len(s1), len(s2)
	if m > n {
		return false
	}

	var count [26]int
	for i := range s1 {
		count[s1[i] - 'a']--
	}

	left := 0
	for i := range s2 {
		char := s2[i] - 'a'
		count[char]++
		for count[char] > 0 {
			count[s2[left] - 'a']--
			left++
		}

		if i - left + 1 == m {
			return true
		}
	}
	return false
}

func main() {
	s1, s2 := "ab", "eidbcaooo"
	res := checkInclusion(s1, s2)
	fmt.Println(res)
}
