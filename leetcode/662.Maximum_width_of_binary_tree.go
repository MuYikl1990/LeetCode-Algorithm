package main

type TreeNode662 struct {
    Val int
    Left *TreeNode662
    Right *TreeNode662
}

func widthOfBinaryTree(root *TreeNode662) int {
	if root == nil {
	    return 0
    }

    var queue []int
	var nodes []*TreeNode662
    queue = append(queue, 1)
    nodes = append(nodes, root)
    res := 1
    for len(nodes) > 0 {
        count := len(nodes)
        for count > 0 {
            cur := nodes[0]
            num := queue[0]
            if cur.Left != nil {
                nodes = append(nodes, cur.Left)
                queue = append(queue, num * 2)
            }
            if cur.Right != nil {
                nodes = append(nodes, cur.Right)
                queue = append(queue, num * 2 + 1)
            }
            nodes = nodes[1:]
            queue = queue[1:]
            count--
        }
        if len(queue) > 1 {
            first, last := queue[0], queue[len(queue) - 1]
            if last - first + 1 > res {
                res = last - first + 1
            }
        }
    }
    return res
}
