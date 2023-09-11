package main

import (
	"fmt"
	"sort"
)

func permuteUnique(nums []int) [][]int {
	sort.Ints(nums)
	var res [][]int
	used := make([]bool, len(nums))
	backtrace47(nums, &res, []int{}, used)
	return res
}

func backtrace47(nums []int, res *[][]int, list []int, used []bool) {
	if len(nums) == len(list) {
		tmp := make([]int, len(list))
		copy(tmp, list)
		*res = append(*res, tmp)
		return
	}

	for i := 0; i < len(nums); i++ {
		if used[i] {
			continue
		}
		if i > 0 && nums[i-1] == nums[i] && !used[i-1] {
			continue
		}
		num := nums[i]
		list = append(list, num)
		used[i] = true
		backtrace47(nums, res, list, used)
		list = list[:len(list)-1]
		used[i] = false
	}
	return
}

func main() {
	nums := []int{1,2,1}
	res := permuteUnique(nums)
	fmt.Println(res)
}
