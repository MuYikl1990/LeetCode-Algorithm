package main

import "fmt"

func convertToTitle(columnNumber int) string {
	var res []rune
	for columnNumber > 0 {
		// 不从0开始
		columnNumber--
		res = append(res, rune('A' + columnNumber % 26))
		columnNumber /= 26
	}
	for i := 0; i < len(res) / 2; i++ {
		res[i], res[len(res)-i-1] = res[len(res)-i-1], res[i]
	}
	return string(res)
}

func main() {
	columnNumber := 52
	res := convertToTitle(columnNumber)
	fmt.Println(res)
}
