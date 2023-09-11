package main

import (
	"fmt"
	"sort"
)

func minMeetingRooms(intervals [][]int) int {
	if len(intervals) == 0 || len(intervals[0]) == 0 {
		return 0
	}

	var begins, ends []int
	for _, interval := range intervals {
		begins = append(begins, interval[0])
		ends = append(ends, interval[1])
	}

	sort.Ints(begins)
	sort.Ints(ends)

	res, end := 0, 0
	for i := 0; i < len(begins); i++ {
		if begins[i] < ends[end] {
			res++
		} else {
			end++
		}
	}
	return res
}

func main() {
	intervals := [][]int{{0, 30}, {5, 10}, {15, 20}, {1, 6}}
	res := minMeetingRooms(intervals)
	fmt.Println(res)
}
