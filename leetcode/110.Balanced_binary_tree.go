package main

func isBalanced(root *TreeNode) bool {
	return depth(root) != -1
}

func depth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := depth(root.Left)
	if left == -1 {
		return -1
	}
	right := depth(root.Right)
	if right == -1 {
		return -1
	}
	if abs(left, right) > 1 {
		return -1
	}
	return max(left, right) + 1
}

func abs(a, b int) int {
	if a >= b {
		return a - b
	}
	return b - a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
