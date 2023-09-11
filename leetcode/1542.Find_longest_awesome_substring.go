package main

import "fmt"

func longestAwesome(s string) int {
	strMap := make(map[int]int)
	state, res := 0, 1
	strMap[state] = -1

	for i := 0; i < len(s); i++ {
		cur := s[i] - '0'
		state ^= 1 << cur
		for k := 0; k <= 9; k++ {
			odd := state ^ 1 << k
			if pos, ok := strMap[odd]; ok {
				res = max1542(res, i - pos)
			}
		}

		if pos, ok := strMap[state]; ok {
			res = max1542(res, i - pos)
		} else {
			strMap[state] = i
		}
	}
	return res
}

func max1542(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	s := "3242415"
	res := longestAwesome(s)
	fmt.Println(res)
}
