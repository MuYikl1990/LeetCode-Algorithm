package main

import (
	"fmt"
	"sort"
)

func wiggleSort(nums []int) {
	if len(nums) < 2 {
		return
	}

	sort.Ints(nums)
	a, b := make([]int, (len(nums) + 1) / 2), make([]int, len(nums) / 2)
	l, r := (len(nums) - 1) / 2, len(nums) - 1
	for i := 0; i < len(a); i++ {
		a[i] = nums[l]
		l--
	}
	for j := 0; j < len(b); j++ {
		b[j] = nums[r]
		r--
	}

	left, right := 0, 0
	for i := 0; i < len(nums); i++ {
		if i % 2 == 0 {
			nums[i] = a[left]
			left++
		} else {
			nums[i] = b[right]
			right++
		}
	}
	return
}

func main() {
	nums := []int{1,3,2,2,3,1}
	wiggleSort(nums)
	fmt.Println(nums)
}
