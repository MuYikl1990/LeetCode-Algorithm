package main

import "fmt"

func findLengthOfShortestSubarray(arr []int) int {
	left, right := 1, len(arr) - 1
	for left < len(arr) && arr[left - 1] <= arr[left] {
		left++
	}
	for right >= 1 && arr[right - 1] <= arr[right] {
		right--
	}
	if left > right {
		return 0
	}

	res := min1574(right, len(arr) - left)
	for i, j := 0, right; i < left && j < len(arr); {
		if arr[i] <= arr[j] {
			res = min1574(res, j - i - 1)
			i++
		} else {
			j++
		}
	}
	return res
}

func min1574(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	arr := []int{10,13,17,21,15,15,9,17,22,22,13}
	res := findLengthOfShortestSubarray(arr)
	fmt.Println(res)
}
