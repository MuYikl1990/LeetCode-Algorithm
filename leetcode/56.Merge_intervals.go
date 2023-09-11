package main

import (
	"fmt"
	"sort"
)

func merge(intervals [][]int) [][]int {
	sort.SliceStable(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0] || (intervals[i][0] == intervals[j][0] && intervals[i][1] <= intervals[j][1])
	})

	var res [][]int
	for i := 0; i < len(intervals); {
		end := intervals[i][1]
		j := i + 1
		for j < len(intervals) && intervals[j][0] <= end {
			end = Max(end, intervals[j][1])
			j++
		}
		res = append(res, []int{intervals[i][0], end})
		i = j
	}
	return res
}

func main() {
	intervals1 := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	intervals2 := [][]int{{1, 4}, {4, 5}}
	res1, res2 := merge(intervals1), merge(intervals2)
	fmt.Println(res1, res2)
}
