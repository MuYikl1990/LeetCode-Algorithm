package main

import "fmt"

func singleNumberIII(nums []int) []int {
	xor := 0
	for _, num := range nums {
		xor ^= num
	}

	res := []int{0, 0}
	n := xor & -xor
	for _, num := range nums {
		if n & num == 0 {
			res[0] ^= num
		} else {
			res[1] ^= num
		}
	}
	return res
}

func main() {
	nums := []int{1,2,1,3,2,5}
	res := singleNumberIII(nums)
	fmt.Println(res)
}
