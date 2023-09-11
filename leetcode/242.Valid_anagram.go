package main

import "fmt"

func isAnagram(s string, t string) bool {
	sMap := make(map[byte]int)
	for i := range s {
		sMap[s[i]]++
	}
	for i := range t {
		sMap[t[i]]--
	}
	for _, val := range sMap {
		if val != 0 {
			return false
		}
	}
	return true
}

func main() {
	s, t := "rat", "car"
	res := isAnagram(s, t)
	fmt.Println(res)
}
