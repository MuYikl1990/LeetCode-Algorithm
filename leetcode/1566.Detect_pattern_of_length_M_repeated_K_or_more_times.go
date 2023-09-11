package main

import "fmt"

func containsPattern(arr []int, m int, k int) bool {
	n := len(arr)
	if n < m * k {
		return false
	}
	for i := 0; i <= n - m * k; i++ {
		j := i + m
		for ; j < i + m * k; j++ {
			if arr[j] != arr[j - m] {
				break
			}
		}
		if j == i + m * k {
			return true
		}
	}
	return false
}

func main() {
	arr, m, k := []int{1,2,1,2,1,3}, 2, 3
	res := containsPattern(arr, m, k)
	fmt.Println(res)
}
