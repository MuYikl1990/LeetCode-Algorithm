package main

import "fmt"

func isRectangleOverlap(rec1 []int, rec2 []int) bool {
	return !(rec1[2] <= rec2[0] || rec2[2] <= rec1[0]) && !(rec1[3] <= rec2[1] || rec2[3] <= rec1[1])
}

func main() {
	rec1, rec2 := []int{5,15,8,18}, []int{0,3,7,9}
	res := isRectangleOverlap(rec1, rec2)
	fmt.Println(res)
}
