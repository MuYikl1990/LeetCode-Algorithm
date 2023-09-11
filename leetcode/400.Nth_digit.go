package main

import (
	"fmt"
	"math"
	"strconv"
)

func findNthDigit(n int) int {
	length, digits := 1, 9
	for n > length * digits {
		n -= length * digits
		length++
		digits *= 10
	}
	n -= 1
	num := int(math.Pow10(length - 1)) + n / length
	index := n % length
	str := strconv.Itoa(num)[index]
	res := str - '0'
	return int(res)
}

func main() {
	n := 201
	res := findNthDigit(n)
	fmt.Println(res)
}
