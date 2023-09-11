package main

import "fmt"

func longestMountain(arr []int) int {
	res, start := 0, -1
	for i := 1; i < len(arr); i++ {
		if arr[i - 1] < arr[i] {
			if i == 1 || arr[i - 2] >= arr[i - 1] {
				start = i - 1
			}
		} else if arr[i - 1] > arr[i] {
			if start != -1 {
				res = max845(res, i - start + 1)
			}
		} else {
			start = -1
		}
	}
	return res
}

func max845(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	arr := []int{2,1,4,7,3,2,5}
	res := longestMountain(arr)
	fmt.Println(res)
}
