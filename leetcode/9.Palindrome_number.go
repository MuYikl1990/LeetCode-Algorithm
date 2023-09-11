package main

import "fmt"

func isPalindromeNumber(x int) bool {
	if x < 0 || (x > 0 && x % 10 == 0) {
		return false
	}
	res := 0
	for x > res {
		cur := x % 10
		res = res * 10 + cur
		x /= 10
	}
	return x == res || (res / 10) == x
}

func main() {
	x := 10
	res := isPalindromeNumber(x)
	fmt.Println(res)
}
