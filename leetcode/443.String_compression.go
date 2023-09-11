package main

import (
	"fmt"
	"strconv"
)

func compress(chars []byte) int {
	left, i := 0, 0
	for left < len(chars) {
		right := left
		for right < len(chars) && chars[left] == chars[right] {
			right++
		}
		chars[i] = chars[left]
		i++
		if right - left > 1 {
			pos := strconv.Itoa(right - left)
			for j := range pos {
				chars[i] = pos[j]
				i++
			}
		}
		left = right
	}
	return i
}

func main() {
	chars := []byte{'a','a','b','b','c','c','c'}
	res := compress(chars)
	fmt.Println(res)
}
