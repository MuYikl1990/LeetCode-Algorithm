package main

import (
	"fmt"
	"math"
)

func integerBreak(n int) int {
	if n < 4 {
		return n - 1
	}

	quo := float64(n / 3)
	mod := n % 3
	if mod == 0 {
		return int(math.Pow(3, quo))
	} else if mod == 1 {
		return int(math.Pow(3, quo - 1) * 4)
	} else {
		return int(math.Pow(3, quo) * 2)
	}
}

func main() {
	n := 11
	res := integerBreak(n)
	fmt.Println(res)
}
