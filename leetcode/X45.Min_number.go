package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func minNumber(nums []int) string {
	res := make([]string, len(nums))
	sort.Slice(nums, func(i, j int) bool {
		return strconv.Itoa(nums[i]) + strconv.Itoa(nums[j]) < strconv.Itoa(nums[j]) + strconv.Itoa(nums[i])
	})

	for i := range nums {
		res = append(res, strconv.Itoa(nums[i]))
	}
	return strings.Join(res, "")

	//res := ""
	//sort.Slice(nums, func(i, j int) bool {
	//	return strconv.Itoa(nums[i]) + strconv.Itoa(nums[j]) < strconv.Itoa(nums[j]) + strconv.Itoa(nums[i])
	//})
	//
	//for i := range nums {
	//	res += strconv.Itoa(nums[i])
	//}
	//return res
}

func main() {
	nums := []int{3,30,34,5,9}
	res := minNumber(nums)
	fmt.Println(res)
}
