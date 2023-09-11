package main

import (
	"fmt"
	"sort"
)

func eraseOverlapIntervals(intervals [][]int) int {
	length := len(intervals)
	if length == 0 {
		return 0
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})

	res, end := 1, intervals[0][1]
	for i := 1; i < length; i++ {
		if intervals[i][0] >= end {
			res++
			end = intervals[i][1]
		}
	}
	return length - res
}

func main() {
	intervals := [][]int{{1, 2}, {2, 3}, {3, 4}, {1, 3}, {1, 2}}
	res := eraseOverlapIntervals(intervals)
	fmt.Println(res)
}
