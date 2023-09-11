package main

import "fmt"

func canThreePartsEqualSum(arr []int) bool {
	target := 0
	for _, num := range arr {
		target += num
	}
	if target % 3 != 0 {
		return false
	}
	target /= 3
	count, sum := 0, 0
	for _, num := range arr {
		sum += num
		if sum == target {
			count++
			sum = 0
		}
		if count == 3 {
			return true
		}
	}
	return false
}

func main() {
	arr := []int{0,2,1,-6,6,-7,9,1,2,0,1}
	res := canThreePartsEqualSum(arr)
	fmt.Println(res)
}
