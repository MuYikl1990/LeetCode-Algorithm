package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	var res [][]int
	if len(nums) < 3 {
		return res
	}

	sort.Ints(nums)
	for i := range nums {
		if nums[i] > 0 {
			return res
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		left, right := i + 1, len(nums) - 1
		for left < right {
			if nums[left] + nums[i] + nums[right] == 0 {
				res = append(res, []int{nums[left], nums[i], nums[right]})
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				left++
				right--
			} else if nums[left] + nums[i] + nums[right] < 0 {
				left++
			} else {
				right--
			}
		}
	}
	return res
}

func main() {
	nums := []int{-1,0,1,2,-1,-4}
	res := threeSum(nums)
	fmt.Println(res)
}
