package main

import "fmt"

func candy(ratings []int) int {
	res := 1
	inc, dec, pre := 1, 0, 1
	for i := 1; i < len(ratings); i++ {
		if ratings[i] >= ratings[i - 1] {
			if ratings[i] == ratings[i - 1] {
				pre = 1
			} else {
				pre++
			}
			dec = 0
			res += pre
			inc = pre
		} else {
			dec++
			if inc == dec {
				dec++
			}
			res += dec
			pre = 1
		}
	}
	return res
}

func main() {
	ratings := []int{1,0,2}
	res := candy(ratings)
	fmt.Println(res)
}
