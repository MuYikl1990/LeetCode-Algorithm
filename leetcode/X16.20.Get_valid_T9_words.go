package main

import (
	"fmt"
	"strconv"
)

func getValidT9Words(num string, words []string) []string {
	var res []string
	table := [26]int{2, 2, 2, 3, 3, 3, 4, 4, 4, 5, 5, 5, 6, 6, 6, 7, 7, 7, 7, 8, 8, 8, 9, 9, 9, 9}
	for _, word := range words {
		target := ""
		for _, char := range word {
			target += strconv.Itoa(table[char - 'a'])
		}
		if target == num {
			res = append(res, word)
		}
	}
	return res
}

func main() {
	num, words := "8733", []string{"tree", "used"}
	res := getValidT9Words(num, words)
	fmt.Println(res)
}
