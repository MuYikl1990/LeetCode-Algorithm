package main

import (
	"fmt"
	"math"
)

func reverseInteger(x int) int {
	res, index := 0, 10
	for x != 0 {
		cur := x % 10
		if res > math.MaxInt32 / 10 || res < math.MinInt32 / 10 {
			return 0
		}
		x /= 10
		res = res * index + cur
	}
	return res
}

func main() {
	x := -120
	res := reverseInteger(x)
	fmt.Println(res)
}
