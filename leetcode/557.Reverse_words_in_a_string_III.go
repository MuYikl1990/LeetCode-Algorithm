package main

import (
	"fmt"
	"strings"
)

func reverseWordsIII(s string) string {
	list := strings.Split(s, " ")
	for i := range list {
		list[i] = stringSwap(list[i])
	}
	return strings.Join(list, " ")
}

func stringSwap(s string) string {
	stringByte := []byte(s)
	n := len(stringByte)
	for i, j := 0, n - 1; i < j; i, j = i + 1, j - 1 {
		stringByte[i], stringByte[j] = stringByte[j], stringByte[i]
	}
	return string(stringByte)
}

func main() {
	s := "Let's take LeetCode contest"
	res := reverseWordsIII(s)
	fmt.Println(res == "s'teL ekat edoCteeL tsetnoc")
}
