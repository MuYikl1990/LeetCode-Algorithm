package main

import "fmt"

func maxChunksToSorted(arr []int) int {
	var stack []int
	for i := range arr {
		if len(stack) > 0 && stack[len(stack) - 1] > arr[i] {
			head := stack[len(stack) - 1]
			for len(stack) > 0 && stack[len(stack) - 1] > arr[i] {
				stack = stack[:len(stack) - 1]
			}
			stack = append(stack, head)
		} else {
			stack = append(stack, arr[i])
		}
	}
	return len(stack)
}

func main() {
	arr := []int{4,2,2,1,1}
	res := maxChunksToSorted(arr)
	fmt.Println(res)
}
