package main

import "fmt"

func singleNumberII(nums []int) int {
	twos, ones := 0, 0
	for _, num := range nums {
		ones = ones ^ num & ^twos
		twos = twos ^ num & ^ones
	}
	return ones
}

func main() {
	nums := []int{0,1,0,1,0,1,99}
	res := singleNumberII(nums)
	fmt.Println(res)
}
