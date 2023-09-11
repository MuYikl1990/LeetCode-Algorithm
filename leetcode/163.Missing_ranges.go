package main

import (
	"fmt"
	"strconv"
)

func findMissingRanges(nums []int, lower, upper int) []string {
	var res []string
	pre := lower
	for i, num := range nums {
		if i == 0 && num > pre {
			res = append(res, numToString(pre, num - 1))
		} else if num - pre > 1 {
			res = append(res, numToString(pre + 1, num - 1))
		}
		pre = num
		if i == len(nums) - 1 && num < upper {
			res = append(res, numToString(num + 1, upper))
		}
	}
	return res
}

func numToString(lower, upper int) string {
	if lower == upper {
		return strconv.Itoa(lower)
	}
	return strconv.Itoa(lower) + "->" + strconv.Itoa(upper)
}

func main() {
	nums, lower, upper := []int{1, 2, 3, 50, 98, 99}, -2, 99
	res := findMissingRanges(nums, lower, upper)
	fmt.Println(res)
}
