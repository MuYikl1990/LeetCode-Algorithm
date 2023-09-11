package main

func btreeGameWinningMove(root *TreeNode, n int, x int) bool {
	half := n / 2
	left, right := 0, 0
	dfs1145(root, &left, &right, x)
	parent := n - left - right - 1
	if parent > half || left > half || right > half {
		return true
	}
	return false
}

func dfs1145(root *TreeNode, left, right *int, target int) int {
	if root == nil {
		return 0
	}

	x := dfs1145(root.Left, left, right, target)
	y := dfs1145(root.Right, left, right, target)

	if root.Val == target {
		*left = x
		*right = y
	}
	return x + y + 1
}
