package main

import "fmt"

func peakIndexInMountainArray(arr []int) int {
	left, right := 0, len(arr) - 1
	for left < right {
		mid := left + (right - left) / 2
		if arr[mid] < arr[mid + 1] {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left
}

func main() {
	arr := []int{0,10,5,2}
	res := peakIndexInMountainArray(arr)
	fmt.Println(res)
}
