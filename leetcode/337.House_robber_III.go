package main

func robIII(root *TreeNode) int {
	res := dfs337(root)
	return max337(res[0], res[1])
}

func dfs337(root *TreeNode) [2]int {
	if root == nil {
		return [2]int{0, 0}
	}

	left := dfs337(root.Left)
	right := dfs337(root.Right)

	var res [2]int
	res[0] = max337(left[0], left[1]) + max337(right[0], right[1])
	res[1] = left[0] + right[0] + root.Val
	return res
}

func max337(a, b int) int {
	if a > b {
		return a
	}
	return b
}
