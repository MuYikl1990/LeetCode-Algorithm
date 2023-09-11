package main

import (
	"fmt"
	"sort"
)

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)

	var res [][]int
	backtrace40(candidates, 0, []int{}, &res, 0, target)
	return res
}

func backtrace40(candidates []int, i int, list []int, res *[][]int, sum, target int) {
	if sum > target {
		return
	}
	if sum == target {
		temp := make([]int, len(list))
		copy(temp, list)
		*res = append(*res, temp)
		return
	}

	for pos := i; pos < len(candidates); pos++ {
		if pos > i && candidates[pos - 1] == candidates[pos] {
			continue
		}
		cur := candidates[pos]
		list = append(list, cur)
		backtrace40(candidates, pos + 1, list, res, sum + cur, target)
		list = list[:len(list)-1]
	}
}

func main() {
	candidates, target := []int{10,1,2,7,6,1,5,1}, 8
	res := combinationSum2(candidates, target)
	fmt.Println(res)
}
