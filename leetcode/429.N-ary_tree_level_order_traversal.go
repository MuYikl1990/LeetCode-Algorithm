package main

type Node429 struct {
    Val int
    Children []*Node429
}


func NLevelOrder(root *Node429) [][]int {
	var res [][]int
	if root == nil {
		return res
	}
	var stack []*Node429
	stack = append(stack, root)
	for len(stack) > 0 {
		length := len(stack)
		var tmp []int
		for i := 0; i < length; i++ {
			cur := stack[i]
			tmp = append(tmp, cur.Val)
			if len(cur.Children) > 0 {
				stack = append(stack, cur.Children...)
			}
		}
		stack = stack[length:]
		res = append(res, tmp)
	}
	return res
}
