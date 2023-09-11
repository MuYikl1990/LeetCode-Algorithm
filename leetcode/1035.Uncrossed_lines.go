package main

import "fmt"

func maxUncrossedLines(nums1 []int, nums2 []int) int {
	m, n := len(nums1), len(nums2)
	dp := make([][]int, m + 1)
	for i := 0; i < m + 1; i++ {
		dp[i] = make([]int, n + 1)
	}

	for i := 1; i < m + 1; i++ {
		for j := 1; j < n + 1; j++ {
			if nums1[i - 1] == nums2[j - 1] {
				dp[i][j] = max1035(dp[i][j], dp[i - 1][j - 1] + 1)
			} else {
				dp[i][j] = max1035(dp[i - 1][j], dp[i][j - 1])
			}
		}
	}
	return dp[m][n]
}

func max1035(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	nums1, nums2 := []int{2,5,1,2,5}, []int{10,5,2,1,5,2}
	res := maxUncrossedLines(nums1, nums2)
	fmt.Println(res)
}
