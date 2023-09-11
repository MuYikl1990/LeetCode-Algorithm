package main

import (
	"fmt"
	"sort"
)

func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	var res [][]int
	dfs90(nums, 0, []int{}, &res)
	return res
}

func dfs90(nums []int, start int, ans []int, res *[][]int) {
	tmp := make([]int, len(ans))
	copy(tmp, ans)
	*res = append(*res, tmp)

	for i := start; i < len(nums); i++ {
		if i > start && nums[i] == nums[i - 1] {
			continue
		}
		ans = append(ans, nums[i])
		dfs90(nums, i + 1, ans, res)
		ans = ans[:len(ans) - 1]
	}
}

func main() {
	nums := []int{1,2,2}
	res := subsetsWithDup(nums)
	fmt.Println(res)
}
