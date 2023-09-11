package main

import "fmt"

func rob(nums []int) int {
	pre, cur := 0, 0
	for _, num := range nums {
		pre, cur = cur, max198(pre + num, cur)
	}
	return cur
}

func max198(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	nums := []int{2,7,9,3,1}
	res := rob(nums)
	fmt.Println(res)
}
