package main

import "fmt"

func longestWPI(hours []int) int {
	pos := make(map[int]int)
	res, sum := 0, 0

	for i := range hours {
		if hours[i] > 8 {
			sum++
		} else {
			sum--
		}

		if sum > 0 {
			res = i + 1
		} else {
			if val, ok := pos[sum - 1]; ok {
				res = max1124(res, i - val)
			}
		}

		if _, ok := pos[sum]; !ok {
			pos[sum] = i
		}
	}
	return res
}

func max1124(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	hours := []int{9,9,6,0,6,6,9}
	res := longestWPI(hours)
	fmt.Println(res)
}
