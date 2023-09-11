package main

import "fmt"

// 异或运算: x ^ 0 = x, x ^ 1 = ^x, 与运算: x & 0 = 0, x & 1 = x
func findSingleNumberII(nums []int) int {
	ones, twos := 0, 0
	for _, num := range nums {
		ones = ones ^ num & ^twos
		twos = twos ^ num & ^ones
	}
	return ones
}

func main() {
	nums := []int{9,1,7,9,7,9,7}
	res := findSingleNumberII(nums)
	fmt.Println(res)
}
