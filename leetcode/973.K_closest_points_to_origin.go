package main

import (
	"fmt"
)

func kClosest(points [][]int, k int) [][]int {
	if len(points) == k {
		return points
	}
	var distance func([]int) int
	distance = func(point []int) int {
		return point[0] * point[0] + point[1] * point[1]
	}
	var quickSelect func([][]int, int, int, int)
	quickSelect = func(points [][]int, k int, start, end int) {
		left, right := start, end
		pivot := distance(points[start])
		for left <= right {
			if distance(points[left]) <= pivot {
				left++
				continue
			}
			if distance(points[right]) > pivot {
				right--
				continue
			}
			points[left], points[right] = points[right], points[left]
			left++
			right--
		}
		points[start], points[right] = points[right], points[start]
		if right == k {
			return
		} else if right < k {
			quickSelect(points, k, right + 1, end)
		} else {
			quickSelect(points, k, start, right - 1)
		}
	}
	quickSelect(points, k, 0, len(points) - 1)
	return points[:k]
}

func main() {
	points, k := [][]int{{3,3}, {5,-1}, {-2,4}}, 2
	res := kClosest(points, k)
	fmt.Println(res)
}
