package main

import (
	"fmt"
	"math"
)

func divide(dividend int, divisor int) int {
	isNeg := false
	if dividend > 0 && divisor < 0 || dividend < 0 && divisor > 0 {
		isNeg = true
	}
	if dividend < 0 {
		dividend = -dividend
	}
	if divisor < 0 {
		divisor = -divisor
	}
	left, right := 0, dividend
	for left <= right {
		mid := left + (right - left) >> 1
		if mul(mid, divisor) <= dividend {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	if isNeg {
		right = -right
	}
	if right > math.MaxInt32 {
		return math.MaxInt32
	}
	if right < math.MinInt32 {
		return math.MinInt32
	}
	return right
}

// 倍增乘法
func mul(a, k int) int {
	ans := 0
	for k > 0 {
		if (k & 1) == 1 {
			ans += a
		}
		k >>= 1
		a += a
	}
	return ans
}

func main() {
	dividend, divisor := 10, 3
	res := divide(dividend, divisor)
	fmt.Println(res)
}
