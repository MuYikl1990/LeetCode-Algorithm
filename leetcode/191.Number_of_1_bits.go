package main

import "fmt"

func hammingWeight(num uint32) int {
	res := 0
	for num != 0 {
		res++
		num &= num - 1
	}
	return res
}

func main() {
	num := uint32(11)
	res := hammingWeight(num)
	fmt.Println(res)
}
