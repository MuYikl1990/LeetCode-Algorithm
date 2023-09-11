package main

import "fmt"

func validUtf8(data []int) bool {
	index := 0

	for _, num := range data {
		if index == 0 {
			if num >> 5 == 0b110 {
				index = 1
			} else if num >> 4 == 0b1110 {
				index = 2
			} else if num >> 3 == 0b11110 {
				index = 3
			} else if num >> 7 == 0b1 {
				return false
			}
		} else {
			if num >> 6 != 0b10 {
				return false
			}
			index--
		}
	}
	return index == 0
}

func main() {
	data := []int{197,130,1}
	res := validUtf8(data)
	fmt.Println(res)
}
