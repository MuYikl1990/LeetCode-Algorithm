package main

import "fmt"

func subsets(nums []int) [][]int {
	var res [][]int
	var dfs func(int, []int)
	dfs = func(index int, list []int) {
		tmp := make([]int, len(list))
		copy(tmp, list)
		res = append(res, tmp)
		for i := index; i < len(nums); i++ {
			list = append(list, nums[i])
			dfs(i + 1, list)
			list = list[:len(list) - 1]
		}
		return
	}
	dfs(0, []int{})
	return res
}

func main() {
	nums := []int{1,2,3}
	res := subsets(nums)
	fmt.Println(res)
}
