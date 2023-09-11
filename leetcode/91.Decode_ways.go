package main

import "fmt"

func numDecodings(s string) int {
	if s[0] == '0' {
		return 0
	}
	f2, f1, f := 1, 1, 0
	for i := 1; i < len(s); i++ {
		f = f1
		if s[i] == '0' {
			if s[i-1] == '1' || s[i-1] == '2' {
				f = f2
			} else {
				return 0
			}
		} else {
			if s[i-1] == '1' || (s[i-1] == '2' && (s[i] - '0') <= 6) {
				f = f2 + f1
			}
		}
		f2, f1 = f1, f
	}
	return f1
}

func main() {
	s := "1201"
	res := numDecodings(s)
	fmt.Println(res)
}
