package main

import "fmt"

func getLeastNumbers(arr []int, k int) []int {
	if len(arr) == k {
		return arr
	}
	var partition func([]int, int, int, int)
	partition = func(arr []int, k int, start, end int) {
		pivot, left, right := arr[start], start, end
		for left <= right {
			if arr[left] <= pivot {
				left++
				continue
			}
			if arr[right] > pivot {
				right--
				continue
			}
			arr[left], arr[right] = arr[right], arr[left]
			left++
			right--
		}
		arr[start], arr[right] = arr[right], arr[start]
		if right == k {
			return
		} else if right < k {
			partition(arr, k, right + 1, end)
		} else {
			partition(arr, k, start, right - 1)
		}
	}
	partition(arr, k, 0, len(arr) - 1)
	return arr[:k]
}

func main() {
	arr, k := []int{0,1,2,1}, 1
	res := getLeastNumbers(arr, k)
	fmt.Println(res)
}
