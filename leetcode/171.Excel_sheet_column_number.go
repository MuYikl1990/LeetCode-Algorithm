package main

import (
	"fmt"
)

func titleToNumber(columnTitle string) int {
	res := 0
	for i := range columnTitle {
		res = res * 26 + int(columnTitle[i] - 'A' + 1)
	}
	return res
}

func main() {
	columnTitle := "ZY"
	res := titleToNumber(columnTitle)
	fmt.Println(res)
}
