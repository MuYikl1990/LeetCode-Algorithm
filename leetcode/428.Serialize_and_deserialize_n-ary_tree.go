package main

import (
	"strconv"
	"strings"
)

type TreeNodeN struct {
	Val      int
	Children []*TreeNodeN
}

func serializeN(root *TreeNodeN) string {
	if root == nil {
		return ""
	}

	var queue []*TreeNodeN
	queue = append(queue, root)
	res := []string{strconv.Itoa(root.Val), "#"}

	for len(queue) > 0 {
		node := queue[0]
		if node != nil {
			if len(node.Children) > 0 {
				for _, child := range node.Children {
					queue = append(queue, child)
					res = append(res, strconv.Itoa(child.Val))
				}
			}
		}
		res = append(res, "#")
		queue = queue[1:]
	}

	if len(res) > 0 && res[len(res) - 1] == "#" {
		res = res[:len(res) - 1]
	}
	return strings.Join(res, ",")
}

func deserializeN(data string) *TreeNodeN {
	if data == "" {
		return nil
	}

	ans := strings.Split(data, ",")
	num, _ := strconv.Atoi(ans[0])
	root := &TreeNodeN{Val: num}
	tmp := root
	queue := []*TreeNodeN{root}
	for i := 1; i < len(ans); i++ {
		if ans[i] == "#" {
			tmp = queue[0]
			queue = queue[1:]
		} else {
			num, _ = strconv.Atoi(ans[i])
			node := &TreeNodeN{Val: num}
			tmp.Children = append(tmp.Children, node)
			queue = append(queue, node)
		}
	}
	return root
}

