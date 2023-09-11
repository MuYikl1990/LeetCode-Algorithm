package main

import "fmt"

func computeArea(ax1 int, ay1 int, ax2 int, ay2 int, bx1 int, by1 int, bx2 int, by2 int) int {
	x := max223(0, min223(ax2, bx2) - max223(ax1, bx1))
	y := max223(0, min223(ay2, by2) - max223(ay1, by1))
	return (ax2 - ax1) * (ay2 - ay1) + (bx2 - bx1) * (by2 - by1) - x * y
}

func max223(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min223(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	ax1, ay1, ax2, ay2, bx1, by1, bx2, by2 := -3, 0, 3, 4, 0, -1, 9, 2
	res := computeArea(ax1, ay1, ax2, ay2, bx1, by1, bx2, by2)
	fmt.Println(res)
}
