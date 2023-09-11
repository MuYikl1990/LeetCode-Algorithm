package main

import (
	"fmt"
	"sort"
)

func sortByBits(arr []int) []int {
	arrMap := make(map[int]int)
	for i := range arr {
		num, count := arr[i], 0
		for num != 0 {
			count++
			num &= num - 1
		}
		arrMap[arr[i]] = count
	}
	sort.Slice(arr, func(i, j int) bool {
		return arrMap[arr[i]] < arrMap[arr[j]] || (arrMap[arr[i]] == arrMap[arr[j]] && arr[i] <= arr[j])
	})
	return arr
}

func main() {
	arr := []int{0,1,2,3,4,5,6,7,8}
	res := sortByBits(arr)
	fmt.Println(res)
}
