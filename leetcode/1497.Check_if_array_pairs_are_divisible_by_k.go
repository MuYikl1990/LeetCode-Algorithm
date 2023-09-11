package main

import "fmt"

func canArrange(arr []int, k int) bool {
	pair := make(map[int]int)
	for i := range arr {
		pair[(arr[i] % k + k) % k]++
	}
	for j := 1; j <= k / 2; j++ {
		if pair[j] != pair[k - j] {
			return false
		}
	}
	return pair[0] % 2 == 0
}

func main() {
	arr, k := []int{1,2,3,4,5,10,6,7,8,9}, 5
	res := canArrange(arr, k)
	fmt.Println(res)
}
