package main

import "fmt"

func minEatingSpeed(piles []int, h int) int {
	left, right := 1, -1
	for i := range piles {
		if piles[i] > right {
			right = piles[i]
		}
	}

	for left < right {
		mid := left + (right - left) / 2
		if isEat(piles, mid, h) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

func isEat(piles []int, k, h int) bool {
	time := 0
	for _, pile := range piles {
		time += (pile + k - 1) / k
		if time > h {
			return false
		}
	}
	return true
}

func main() {
	piles, h := []int{30,11,23,4,20}, 6
	res := minEatingSpeed(piles, h)
	fmt.Println(res)
}
