package main

import "fmt"

func validateStackSequences(pushed []int, popped []int) bool {
	j, size := 0, 0
	for i := 0; i < len(pushed); i++ {
		pushed[size] = pushed[i]
		size++
		for size != 0 && popped[j] == pushed[size - 1] {
			j++
			size--
		}
	}
	return size == 0
}

func main() {
	pushed, popped := []int{1, 2, 3, 4, 5}, []int{4, 5, 3, 1, 2}
	res := validateStackSequences(pushed, popped)
	fmt.Println(res)
}
