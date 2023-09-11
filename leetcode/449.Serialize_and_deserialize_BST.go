package main

import (
	"fmt"
	"strconv"
	"strings"
)

type TreeNode449 struct {
    Val int
    Left *TreeNode449
    Right *TreeNode449
}

type Codec449 struct {

}

func Constructor449() Codec449 {
	return Codec449{}
}

// Serializes a tree to a single string.
func (this *Codec449) serialize(root *TreeNode449) string {
	//var res string
	//dfs449(root, &res)
	//return res
	if root == nil {
		return "#"
	}

	return strconv.Itoa(root.Val) + "," + this.serialize(root.Left) + "," + this.serialize(root.Right)
}

func dfs449(root *TreeNode449, res *string) {
	if root == nil {
		return
	}

	*res += strconv.Itoa(root.Val) + " "
	dfs449(root.Left, res)
	dfs449(root.Right, res)
}

// Deserializes your encoded data to tree.
func (this *Codec449) deserialize(data string) *TreeNode449 {
	list := strings.Split(data, ",")
	//fmt.Println(list)
	//return helper(list, 0, len(list) - 1)
	return df(&list)
}

func df(list *[]string) *TreeNode449 {
	cur := (*list)[0]
	*list = (*list)[1:]

	if cur == "#" {
		return nil
	}
	val, _ := strconv.Atoi(cur)
	root := &TreeNode449{Val: val}
	root.Left = df(list)
	root.Right = df(list)
	return root
}

func helper(list []string, start, end int) *TreeNode449 {
	if start > end {
		return nil
	}

	cur := start
	str := list[start]
	num, _ := strconv.Atoi(str)
	root := &TreeNode449{Val: num}
	start += 1
	for start < end && list[start] < str {
		start++
	}
	root.Left = helper(list, cur + 1, start - 1)
	root.Right = helper(list, start, end)
	return root
}

func main() {
	root := &TreeNode449{Val: 2, Left: &TreeNode449{Val: 1}, Right: &TreeNode449{Val: 3}}
	ser := Constructor449()
	deser := Constructor449()
	tree := ser.serialize(root)
	fmt.Println(tree)
	res := deser.deserialize(tree)
	fmt.Println(res)
}
