package main

import (
	"fmt"
	"math"
)

func verifyPostorder(postorder []int) bool {
	var stack []int
	peek := math.MaxInt

	for i := len(postorder) - 1; i >= 0; i-- {
		cur := postorder[i]
		if cur > peek {
			return false
		}

		for len(stack) != 0 && cur < stack[len(stack) - 1] {
			peek = stack[len(stack) - 1]
			stack = stack[:len(stack) - 1]
		}
		stack = append(stack, cur)
	}
	return true
}

func main() {
	postorder := []int{1,2,5,10,6,9,4,3}
	res := verifyPostorder(postorder)
	fmt.Println(res)
}
