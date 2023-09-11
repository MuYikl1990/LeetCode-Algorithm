package main

type TreeNode108 struct {
    Val int
    Left *TreeNode108
    Right *TreeNode108
}

func sortedArrayToBST(nums []int) *TreeNode108 {
	return dfs108(nums, 0, len(nums) - 1)
}

func dfs108(nums []int, left ,right int) *TreeNode108 {
	if left > right {
		return nil
	}
	mid := left + (right - left) / 2
	root := &TreeNode108{Val: nums[mid]}
	root.Left = dfs108(nums, left, mid - 1)
	root.Right = dfs108(nums, mid + 1, right)
	return root
}
