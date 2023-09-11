package main

import "fmt"

func intervalIntersection(firstList [][]int, secondList [][]int) [][]int {
	var res [][]int
	i, j := 0, 0
	for i < len(firstList) && j < len(secondList) {
		start := max986(firstList[i][0], secondList[j][0])
		end := min986(firstList[i][1], secondList[j][1])
		if start <= end {
			res = append(res, []int{start, end})
		}
		if firstList[i][1] < secondList[j][1] {
			i++
		} else {
			j++
		}
	}
	return res
}

func max986(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min986(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	firstList, secondList := [][]int{{0, 2}, {5, 10}, {13, 23}, {24, 25}}, [][]int{{1, 5}, {8, 12}, {15, 24}, {25, 26}}
	res := intervalIntersection(firstList, secondList)
	fmt.Println(res)
}
