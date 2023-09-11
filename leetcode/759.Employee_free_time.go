package main

import (
	"fmt"
	"math"
	"sort"
)

func employeeFreeTime(schedules [][][]int) [][]int {
	var intervals, res [][]int
	end := math.MaxInt
	for _, schedule := range schedules {
		for _, interval := range schedule {
			intervals = append(intervals, interval)
			end = min759(end, interval[1])
		}
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0] || (intervals[i][0] == intervals[j][0] && intervals[i][1] < intervals[j][1])
	})

	for _, interval := range intervals {
		if end < interval[0] {
			res = append(res, []int{end, interval[0]})
		}
		end = max759(end, interval[1])
	}
	return res
}

func min759(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max759(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	schedule := [][][]int{{{1, 3}, {6, 7}}, {{2, 4}},{{2, 5},{9, 12}}}
	res := employeeFreeTime(schedule)
	fmt.Println(res)
}
