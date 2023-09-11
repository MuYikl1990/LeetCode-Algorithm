package main

import (
	"fmt"
	"strconv"
)

func getPermutation(n int, k int) string {
	k--

	fab := make([]int, n)
	fab[0] = 1
	for i := 1; i < n; i++ {
		fab[i] = fab[i - 1] * i
	}

	var list []int
	for i := 0; i < n; i++ {
		list = append(list, i + 1)
	}

	res := ""
	for i := 1; i <= n; i++ {
		index := k / fab[n - i]
		cur := list[index]
		res += strconv.Itoa(cur)
		k = k % fab[n - i]
		list = append(list[:index], list[index + 1:]...)
	}
	return res
}

func main() {
	n, k := 4, 9
	res := getPermutation(n, k)
	fmt.Println(res)
}
