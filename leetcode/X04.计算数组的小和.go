package main

import "fmt"

// 在一个数组中，每一个数左边比当前数小或者相等的数累加起来，叫做这个数组的小和。求一个数组的小和
func getSmallSum(nums []int) int {
	if len(nums) < 2 {
		return 0
	}
	return partNums(nums, 0, len(nums) - 1)
}

func partNums(nums []int, left, right int) int {
	if left == right {
		return 0
	}
	mid := left + (right - left) / 2
	return partNums(nums, left, mid) + partNums(nums, mid + 1, right) + mergeNums(nums, left, mid, right)
}

func mergeNums(nums []int, left, mid, right int) int {
	tmp := make([]int, right - left + 1)
	start, i, j := 0, left, mid + 1
	res := 0
	for i <= mid && j <= right {
		if nums[i] <= nums[j] {
			res += nums[i] * (right - j + 1)
			tmp[start] = nums[i]
			i++
		} else {
			tmp[start] = nums[j]
			j++
		}
		start++
	}

	for i <= mid {
		tmp[start] = nums[i]
		i++
		start++
	}
	for j <= right {
		tmp[start] = nums[j]
		j++
		start++
	}

	for i, start = left, 0; i <= right; i++ {
		nums[i] = tmp[start]
		start++
	}
	return res
}

func main() {
	nums := []int{1,3,4,2,5,6}
	res := getSmallSum(nums)
	fmt.Println(res)
}
