package main

import "fmt"

func permute(nums []int) [][]int {
	var res [][]int
	backtrace46(nums, &res, []int{})
	return res
}

func backtrace46(nums []int, res *[][]int, list []int) {
	if len(nums) == 0 {
		tmp := make([]int, len(list))
		copy(tmp, list)
		*res = append(*res, tmp)
		return
	}

	for i := 0; i < len(nums); i++ {
		cur := nums[i]
		list = append(list, cur)
		nums = append(nums[:i], nums[i+1:]...)
		backtrace46(nums, res, list)
		list = list[:len(list)-1]
		nums = append(nums[:i], append([]int{cur}, nums[i:]...)...)
	}
	return
}

func main() {
	nums := []int{1,2,3}
	res := permute(nums)
	fmt.Println(res)
}
