package main

import "fmt"

func climbStairs(n int) int {
	p, q, res := 0, 0, 1
	for i := 0; i < n; i++ {
		p = q
		q = res
		res = p + q
	}
	return res
}

func main() {
	n := 3
	res := climbStairs(n)
	fmt.Println(res)
}
