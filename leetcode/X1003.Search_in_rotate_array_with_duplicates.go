package main

import "fmt"

func searchRotateArray(arr []int, target int) int {
	length := len(arr)
	left, right := 0, length - 1
	for left <= right {
		if arr[left] == target {
			return left
		}
		mid := left + (right - left) / 2
		if arr[mid] == target {
			right = mid
		} else if arr[0] < arr[mid] {
			if arr[0] <= target && target < arr[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else if arr[0] > arr[mid] {
			if arr[mid] < target && target <= arr[length - 1] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		} else {
			left++
		}
	}
	return -1
}

func main() {
	arr, target := []int{14, 14, 15, 16, 19, 20, 25, 1, 3, 4, 4, 5, 7, 10, 14}, 14
	res := searchRotateArray(arr, target)
	fmt.Println(res)
}
