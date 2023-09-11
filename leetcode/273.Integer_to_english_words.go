package main

import (
	"fmt"
	"strings"
)

var (
	single = []string{"", "One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine"}
	teen = []string{"Ten", "Eleven", "Twelve", "Thirteen", "Fourteen", "Fifteen", "Sixteen", "Seventeen", "Eighteen", "Nineteen"}
	ten = []string{"", "Ten", "Twenty", "Thirty", "Forty", "Fifty", "Sixty", "Seventy", "Eighty", "Ninety"}
	unit = []string{"", "Thousand", "Million", "Billion"}
)

func numberToWords(num int) string {
	if num == 0 {
		return "Zero"
	}

	result := ""
	for i, step := 3, int(1e9); i >= 0; i-- {
		if curNum := num / step; curNum > 0 {
			num -= curNum * step
			toEnglish(curNum, &result)
			result += unit[i] + " "
		}
		step /= 1000
	}
	return strings.TrimSpace(result)
}

func toEnglish(num int, result *string) {
	if num >= 100 {
		*result += single[num / 100] + " Hundred "
		num %= 100
	}
	if num >= 20 {
		*result += ten[num / 10] + " "
		num %= 10
	}
	if 0 < num && num < 10 {
		*result += single[num] + " "
	} else if num >= 10 {
		*result += teen[num - 10] + " "
	}
}

func main() {
	num := 123456
	res := numberToWords(num)
	fmt.Println(res)
}
