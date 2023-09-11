package main

import "fmt"

func romanToInt(s string) int {
	res := 0
	romanMap := map[byte]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}
	for i := 0; i < len(s) - 1; i++ {
		if romanMap[s[i]] < romanMap[s[i + 1]] {
			res -= romanMap[s[i]]
		} else {
			res += romanMap[s[i]]
		}
	}
	return res + romanMap[s[len(s) - 1]]
}

func main() {
	s := "MCMXCIV"
	res := romanToInt(s)
	fmt.Println(res)
}
