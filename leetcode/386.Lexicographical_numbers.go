package main

import "fmt"

func lexicalOrder(n int) []int {
	var res []int
	for i, j := 0, 1; i < n; i++ {
		res = append(res, j)
		if j * 10 <= n {
			j *= 10
		} else {
			for j % 10 == 9 || j + 1 > n {
				j /= 10
			}
			j++
		}
	}
	return res
}

func main() {
	n := 13
	res := lexicalOrder(n)
	fmt.Println(res)
}
