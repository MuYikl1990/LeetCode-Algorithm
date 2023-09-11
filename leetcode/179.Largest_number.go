package main

import (
	"fmt"
	"sort"
	"strconv"
)

func largestNumber(nums []int) string {
	sort.SliceStable(nums, func(i, j int) bool {
		x, y := nums[i], nums[j]
		xIndex, yIndex := 10, 10
		for xIndex <= x {
			xIndex *= 10
		}
		for yIndex <= y {
			yIndex *= 10
		}
		return yIndex * x + y > xIndex * y + x
	})

	if nums[0] == 0 {
		return "0"
	}
	var res []byte
	for _, x := range nums {
		res = append(res, strconv.Itoa(x)...)
	}
	return string(res)
}

func main() {
	nums := []int{3,30,34,5,9}
	res := largestNumber(nums)
	fmt.Println(res)
}
