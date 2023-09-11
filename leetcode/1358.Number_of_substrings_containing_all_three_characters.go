package main

import "fmt"

func numberOfSubstrings(s string) int {
	res := 0
	list := [3]int{}
	left, right, n := 0, 0, len(s)
	for right < n {
		list[s[right] - 'a']++
		right++
		for list[0] > 0 && list[1] > 0 && list[2] > 0 {
			res += n - right + 1
			list[s[left] - 'a']--
			left++
		}
	}
	return res
}

func main() {
	s := "abcabc"
	res := numberOfSubstrings(s)
	fmt.Println(res)
}
