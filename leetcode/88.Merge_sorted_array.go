package main

import "fmt"

func merge88(nums1 []int, m int, nums2 []int, n int)  {
	i, j := m - 1, n - 1
	for i >= 0 && j >= 0 {
		if nums1[i] >= nums2[j] {
			nums1[i + j + 1] = nums1[i]
			i--
		} else {
			nums1[i + j + 1] = nums2[j]
			j--
		}
	}

	if j >= 0 {
		for j >= 0 {
			nums1[j] = nums2[j]
			j--
		}
	}
}

func main() {
	nums1, m, nums2, n := []int{1,2,3,0,0,0}, 3, []int{2,5,6}, 3
	merge88(nums1, m, nums2, n)
	fmt.Println(nums1)
}
