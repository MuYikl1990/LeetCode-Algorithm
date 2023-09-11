package main

import "fmt"

// 动态规划解法
func findLength(nums1 []int, nums2 []int) int {
	m, n := len(nums1), len(nums2)
	dp := make([]int, n + 1)
	res := 0

	for i := 0; i < m; i++ {
		for j := n - 1; j >= 0; j-- {
			if nums1[i] == nums2[j] {
				dp[j + 1] = dp[j] + 1
			} else {
				dp[j + 1] = 0
			}
			if dp[j + 1] > res {
				res = dp[j + 1]
			}
		}
	}
	return res
}

// 滑动窗口解法
func findLengthSlide(nums1 []int, nums2 []int) int {
	if len(nums1) < len(nums2) {
		return findLengthHelper(nums1, nums2)
	}
	return findLengthHelper(nums2, nums1)
}

func findLengthHelper(nums1 []int, nums2 []int) int {
	aLength, bLength := len(nums1), len(nums2)
	res := 0
	aStart, bStart, length := 0, 0, 0
	for i := 0; i < aLength + bLength - 1; i++ {
		if i < aLength {
			aStart = aLength - i - 1
			bStart = 0
			length = i + 1
		} else {
			aStart = 0
			bStart = i - aLength + 1
			length = min718(bLength - bStart, aLength)
		}
		maxLen := maxLength(nums1, nums2, aStart, bStart, length)
		res = max718(res, maxLen)
	}
	return res
}

func maxLength(nums1, nums2 []int, aStart, bStart, length int) int {
	count, ans := 0, 0
	for i := 0; i < length; i++ {
		if nums1[aStart + i] == nums2[bStart + i] {
			count++
			ans = max718(ans, count)
		} else {
			count = 0
		}
	}
	return ans
}

func min718(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max718(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	nums1, nums2 := []int{1,2,3,2,1}, []int{3,2,1,4,7}
	res := findLength(nums1, nums2)
	fmt.Println(res)
	ans := findLengthSlide(nums1, nums2)
	fmt.Println(ans)
}
