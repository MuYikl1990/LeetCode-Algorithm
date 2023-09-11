package main

import "fmt"

func findRepeatedDnaSequences(s string) []string {
	repeat := make(map[string]int)
	var res []string
	for i := 0; i < len(s) - 9; i++ {
		if count, ok := repeat[s[i:i + 10]]; ok {
			repeat[s[i:i + 10]]++
			if count == 1 {
				res = append(res, s[i:i + 10])
			}
		} else {
			repeat[s[i:i + 10]]++
		}
	}
	return res
}

func main() {
	s := "AAAAAAAAAAA"
	res := findRepeatedDnaSequences(s)
	fmt.Println(res)
}
