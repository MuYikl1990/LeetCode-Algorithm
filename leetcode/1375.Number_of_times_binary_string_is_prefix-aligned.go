package main

import "fmt"

func numTimesAllBlue(flips []int) int {
	res := 0
	maxNum := 0
	for i := 0; i < len(flips); i++ {
		maxNum = max1375(maxNum, flips[i])
		if maxNum == i + 1 {
			res++
		}
	}
	return res
}

func max1375(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	flips := []int{3,2,4,1,5}
	res := numTimesAllBlue(flips)
	fmt.Println(res)
}
