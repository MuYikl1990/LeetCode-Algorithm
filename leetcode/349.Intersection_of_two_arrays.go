package main

import "fmt"

func intersection(nums1 []int, nums2 []int) []int {
	var res []int
	numMap := make(map[int]bool)
	for i := range nums1 {
		numMap[nums1[i]] = true
	}

	for j := range nums2 {
		if numMap[nums2[j]] {
			res = append(res, nums2[j])
			numMap[nums2[j]] = false
		}
	}
	return res
}

func main() {
	nums1, nums2 := []int{4,9,5}, []int{9,4,9,8,4}
	res := intersection(nums1, nums2)
	fmt.Println(res)
}
