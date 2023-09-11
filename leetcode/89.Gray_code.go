package main

import "fmt"

func grayCode(n int) []int {
	res, head := []int{0}, 1
	for i := 0; i < n; i++ {
		for j := len(res) - 1; j >= 0; j-- {
			res = append(res, head + res[j])
		}
		head <<= 1
	}
	return res
}

func main() {
	n := 3
	res := grayCode(n)
	fmt.Println(res)
}
