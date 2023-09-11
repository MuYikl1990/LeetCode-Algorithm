package main

import "fmt"

func addI(a int, b int) int {
	for b != 0 {
		a, b = a ^ b, a & b << 1
	}
	return a
}

func main() {
	a, b := 5, -12
	res := addI(a, b)
	fmt.Println(res)
}
