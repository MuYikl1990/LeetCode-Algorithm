package main

import "fmt"

func mySqrt(x int) int {
	if x == 0 || x == 1 {
		return x
	}

	left, right := 1, x / 2
	for left < right {
		mid := left + (right - left + 1) / 2
		if mid > x / mid {
			right = mid - 1
		} else if mid < x / mid {
			left = mid
		} else {
			return mid
		}
	}
	return left
}

func main() {
	x := 678
	res := mySqrt(x)
	fmt.Println(res)
}
