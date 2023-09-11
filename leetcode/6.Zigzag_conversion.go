package main

import (
	"fmt"
	"strings"
)

func convert(s string, numRows int) string {
	if numRows < 2 {
		return s
	}
	res := make([]string, numRows)
	i, flag := 0, 1
	for _, char := range s {
		res[i] += string(char)
		i += flag
		if i == 0 || i == numRows - 1 {
			flag = -flag
		}
	}
	return strings.Join(res, "")
}

func main() {
	s, numRows := "PAYPALISHIRING", 4
	res := convert(s, numRows)
	outcome := "PINALSIGYAHRPI"
	fmt.Println(res == outcome)
}
