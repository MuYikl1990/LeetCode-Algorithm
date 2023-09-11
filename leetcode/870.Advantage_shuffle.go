package main

import (
	"fmt"
	"sort"
)

func advantageCount(nums1 []int, nums2 []int) []int {
	n := len(nums1)
	res, ids := make([]int, n), make([]int, n)
	sort.Ints(nums1)
	for i := range ids {
		ids[i] = i
	}

	sort.Slice(ids, func(i, j int) bool {
		return nums2[ids[i]] <= nums2[ids[j]]
	})

	left, right := 0, n - 1
	for _, num := range nums1 {
		if num > nums2[ids[left]] {
			res[ids[left]] = num
			left++
		} else {
			res[ids[right]] = num
			right--
		}
	}
	return res
}

func main() {
	nums1, nums2 := []int{12,24,8,32}, []int{13,25,32,11}
	res := advantageCount(nums1, nums2)
	fmt.Println(res)
}
