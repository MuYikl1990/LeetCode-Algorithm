package main

import "fmt"

// 快速幂
func myPow(x float64, n int) float64 {
	if x == 0 {
		return x
	}
	if n < 0 {
		x = 1 / x
		n = -n
	}
	res := float64(1)
	for n > 0 {
		if n & 1 == 1 {
			res *= x
		}
		x *= x
		n >>= 1
	}
	return res
}

func main() {
	x, n := 2.10000, 3
	res := myPow(x, n)
	fmt.Println(res)
}
