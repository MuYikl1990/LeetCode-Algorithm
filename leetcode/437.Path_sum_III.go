package main

type TreeNode437 struct {
    Val int
    Left *TreeNode437
    Right *TreeNode437
}

func pathSum(root *TreeNode437, targetSum int) int {
	sumMap := make(map[int]int)
	sumMap[0] = 1
	return dfs437(root, targetSum, &sumMap, 0)
}

func dfs437(root *TreeNode437, targetSum int, sumMap *map[int]int, sum int) int {
	if root == nil {
		return 0
	}

	res := 0
	sum += root.Val
	cnt, _ := (*sumMap)[sum - targetSum]
	res += cnt

	(*sumMap)[sum]++
	left := dfs437(root.Left, targetSum, sumMap, sum)
	right := dfs437(root.Right, targetSum, sumMap, sum)

	res += left + right

	(*sumMap)[sum]--
	return res
}
