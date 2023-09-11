package main

import "fmt"

func intToRoman(num int) string {
	numArr := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	romanArr := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	res := ""
	for i, val := range numArr {
		if num / val > 0 {
			rep := num / val
			num %= val
			for j := 0; j < rep; j++ {
				res += romanArr[i]
			}
		}
	}
	return res
}

func main() {
	num := 1994
	res := intToRoman(num)
	fmt.Println(res)
}
