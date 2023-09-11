package main

import "fmt"

func hanota(A []int, B []int, C []int) []int {
	n := len(A)
	move(n, &A, &B, &C)
	return C
}

func move(n int, A *[]int, B *[]int, C *[]int) {
	if n <= 0 {
		return
	}
	move(n - 1, A, C, B)
	*C = append(*C, (*A)[len(*A) - 1])
	*A = (*A)[:len(*A) - 1]
	move(n - 1, B, A, C)
}

func main() {
	var B, C []int
	A := []int{2,1,0}
	res := hanota(A, B, C)
	fmt.Println(res)
}
