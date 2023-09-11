package main

import "fmt"

func insert(intervals [][]int, newInterval []int) [][]int {
	var res [][]int
	i := 0
	for i < len(intervals) && intervals[i][1] < newInterval[0] {
		res = append(res, intervals[i])
		i++
	}

	for i < len(intervals) && intervals[i][0] <= newInterval[1] {
		newInterval[0] = min57(intervals[i][0], newInterval[0])
		newInterval[1] = max57(intervals[i][1], newInterval[1])
		i++
	}
	res = append(res, newInterval)

	for i < len(intervals) {
		res = append(res, intervals[i])
		i++
	}
	return res
}

func min57(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max57(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	intervals, newInterval := [][]int{{1,2}, {3,5}, {6,7}, {8,10}, {12,16}}, []int{4,8}
	res := insert(intervals, newInterval)
	fmt.Println(res)
}
