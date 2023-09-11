package main

import (
	"strconv"
	"strings"
)

type Codec struct {
}

func Constructor297() Codec {
	return Codec{}
}

func (this *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return "#"
	}

	return strconv.Itoa(root.Val) + "," + this.serialize(root.Left) + "," + this.serialize(root.Right)
}

func (this *Codec) deserialize(data string) *TreeNode {
	list := strings.Split(data, ",")
	return buildTree(&list)
}

func buildTree(list *[]string) *TreeNode {
	if len(*list) == 0 {
		return nil
	}

	cur := (*list)[0]
	*list = (*list)[1:]
	if cur == "#" {
		return nil
	}

	val, _ := strconv.Atoi(cur)
	root := &TreeNode{Val: val}
	root.Left = buildTree(list)
	root.Right = buildTree(list)
	return root
}
