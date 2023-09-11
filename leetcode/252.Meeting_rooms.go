package main

import (
	"fmt"
	"sort"
)

// Given an array of meeting time intervals consisting of start and end times [[s1,e1],[s2,e2],...] (si < ei),
// determine if a person could attend all meetings.
func canAttendMeetings(intervals [][]int) bool {
	sort.SliceStable(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	for i := 0; i < len(intervals) - 1; i++ {
		if intervals[i][1] > intervals[i + 1][0] {
			return false
		}
	}
	return true
}

func main() {
	intervals := [][]int{{0, 3}, {5, 10}, {17, 20}, {10, 17}}
	res := canAttendMeetings(intervals)
	fmt.Println(res)
}
