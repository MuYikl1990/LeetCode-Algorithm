package main

import "fmt"

func dailyTemperatures(temperatures []int) []int {
	res := make([]int, len(temperatures))
	var stack []int
	for i := 0; i < len(temperatures); i++ {
		for len(stack) > 0 && temperatures[stack[len(stack) - 1]] < temperatures[i] {
			index := stack[len(stack) - 1]
			res[stack[len(stack) - 1]] = i - index
			stack = stack[:len(stack) - 1]
		}
		stack = append(stack, i)
	}
	return res
}

func main() {
	temperatures := []int{73,74,75,71,69,72,76,73}
	res := dailyTemperatures(temperatures)
	fmt.Println(res)
}
