package main

import "fmt"

func plusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] == 9 {
			digits[i] = 0
		} else {
			digits[i]++
			return digits
		}
	}
	digits = append([]int{1}, digits...)
	return digits
}

func main() {
	digits := []int{9,9}
	res := plusOne(digits)
	fmt.Println(res)
}
