package main

import (
	"fmt"
	"strconv"
)

func fractionToDecimal(numerator int, denominator int) string {
	num, dem := int64(numerator), int64(denominator)
	res := ""
	if num % dem == 0 {
		res = strconv.Itoa(int(num / dem))
		return res
	}
	if num * dem < 0 {
		res += "-"
	}
	if num < 0 {
		num = -num
	}
	if dem < 0 {
		dem = -dem
	}
	numMap := make(map[int64]int)
	res += strconv.Itoa(int(num / dem)) + "."
	num %= dem
	for num != 0 {
		if pos, ok := numMap[num]; ok {
			res = res[:pos] + "(" + res[pos:] + ")"
			break
		}
		numMap[num] = len(res)
		num *= 10
		res += strconv.Itoa(int(num / dem))
		num %= dem
	}
	return res
}

func main() {
	numerator, denominator := 4, 333
	res := fractionToDecimal(numerator, denominator)
	fmt.Println(res)
}
