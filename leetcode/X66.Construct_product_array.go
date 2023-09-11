package main

import "fmt"

func constructArr(a []int) []int {
	res := make([]int, len(a))
	pre := 1
	for i := range a {
		res[i] = pre
		pre *= a[i]
	}
	post := 1
	for i := len(a) - 1; i >= 0; i-- {
		res[i] *= post
		post *= a[i]
	}
	return res
}

func main() {
	a := []int{1,2,3,4,5}
	res := constructArr(a)
	fmt.Println(res)
}
