package main

import "fmt"

func translateNum(num int) int {
	f1, f2, f := 1, 1, 0
	x, y := 0, num % 10
	for num != 0 {
		num /= 10
		x = num % 10
		if 10 * x + y >= 10 && 10 * x + y <= 25 {
			f = f1 + f2
		} else {
			f = f1
		}
		y = x
		f2 = f1
		f1 = f
	}
	return f1
}

func main() {
	num := 12258
	res := translateNum(num)
	fmt.Println(res)
}
