package main

import (
	"fmt"
	"strconv"
	"strings"
)

func maximumSwap(num int) int {
	nums := strings.Split(strconv.Itoa(num), "")
	pos := make([]int, len(nums))
	maxPos := len(nums) - 1
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] > nums[maxPos] {
			maxPos = i
		}
		pos[i] = maxPos
	}

	for i := 0; i < len(pos); i++ {
		if nums[pos[i]] != nums[i] {
			nums[pos[i]], nums[i] = nums[i], nums[pos[i]]
			break
		}
	}
	res, _ := strconv.Atoi(strings.Join(nums, ""))
	return res
}

func main() {
	num := 2736
	res := maximumSwap(num)
	fmt.Println(res)
}
