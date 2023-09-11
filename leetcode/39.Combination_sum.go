package main

import "fmt"

func combinationSum(candidates []int, target int) [][]int {
	var res [][]int
	backtrace(candidates, target, &res, []int{}, 0, 0)
	return res
}

func backtrace(candidates []int, target int, res *[][]int, queue []int, sum int, start int) {
	if sum == target {
		tmp := make([]int, len(queue))
		copy(tmp, queue)
		*res = append(*res, tmp)
		return
	} else if sum > target {
		return
	}

	for i := start; i < len(candidates); i++ {
		queue = append(queue, candidates[i])
		backtrace(candidates, target, res, queue, sum + candidates[i], i)
		queue = queue[:len(queue)-1]
	}
	return
}

func main() {
	candidates, target := []int{2,3,5}, 8
	res := combinationSum(candidates, target)
	fmt.Println(res)
}
