package main

import "fmt"

func lastRemaining(n int, m int) int {
	res := 0
	for i := 2; i <= n; i++ {
		res = (res + m) % i
	}
	return res
}

func main() {
	n, m := 10, 17
	res := lastRemaining(n, m)
	fmt.Println(res)
}
