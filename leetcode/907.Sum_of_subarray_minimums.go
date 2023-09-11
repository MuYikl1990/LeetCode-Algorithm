package main

import (
	"fmt"
	"math"
)

func sumSubarrayMins(arr []int) int {
	length := len(arr)
	if length == 0 {
		return 0
	}

	res, mod := 0, 1000000007
	var stack []int
	for i := -1; i <= length; i++ {
		for len(stack) != 0 && getValue(arr, length, stack[len(stack) - 1]) > getValue(arr, length, i) {
			cur := stack[len(stack) - 1]
			stack = stack[:len(stack) - 1]
			res = (res + arr[cur] * (cur - stack[len(stack) - 1]) * (i - cur)) % mod
		}
		stack = append(stack, i)
	}
	return res
}

func getValue(arr []int, n, i int) int {
	if i == -1 || i == n {
		return math.MinInt
	}
	return arr[i]
}

func main() {
	arr := []int{3, 1, 2, 4}
	res := sumSubarrayMins(arr)
	fmt.Println(res)
}
