package main

import (
	"fmt"
)

type TreeNode652 struct {
    Val int
    Left *TreeNode652
    Right *TreeNode652
}

func findDuplicateSubtrees(root *TreeNode652) []*TreeNode652 {
	var res []*TreeNode652
	idx, repeat, tmp := 0, map[[3]int]pair{}, map[*TreeNode652]struct{}{}
	dfs652(&idx, repeat, root, tmp)
	for node := range tmp {
		res = append(res, node)
	}
	return res
}

func dfs652(idx *int, repeat map[[3]int]pair, root *TreeNode652, tmp map[*TreeNode652]struct{}) int {
	if root == nil {
		return 0
	}
	cur := [3]int{root.Val, dfs652(idx, repeat, root.Left, tmp), dfs652(idx, repeat, root.Right, tmp)}
	if p, ok := repeat[cur]; ok {
		tmp[p.node] = struct{}{}
		return p.idx
	}
	*idx++
	repeat[cur] = pair{root, *idx}
	return *idx
}

type pair struct {
	node *TreeNode652
	idx int
}

func main() {
	root := &TreeNode652{1, &TreeNode652{2, &TreeNode652{4, nil, nil}, nil}, &TreeNode652{3, &TreeNode652{2, &TreeNode652{4, nil, nil}, nil}, &TreeNode652{4, nil, nil}}}
	res := findDuplicateSubtrees(root)
	fmt.Println(res)
}
