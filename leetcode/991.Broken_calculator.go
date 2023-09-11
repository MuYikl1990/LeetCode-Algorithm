package main

import "fmt"

func brokenCalc(startValue int, target int) int {
	res := 0
	for startValue < target {
		if target % 2 == 1 {
			target += 1
			res++
		}
		target /= 2
		res++
	}
	res += startValue - target
	return res
}

func main() {
	startValue, target := 5, 8
	res := brokenCalc(startValue, target)
	fmt.Println(res)
}
