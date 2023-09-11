package main

import (
	"fmt"
	"sort"
)

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	n := len(nums)
	var res [][]int
	if n < 4 {
		return res
	}

	for a := 0; a < n - 3; a++ {
		if a > 0 && nums[a] == nums[a - 1] {
			continue
		}
		if nums[a] + nums[a + 1] + nums[a + 2] + nums[a + 3] > target {
			break
		}
		if nums[a] + nums[n - 3] + nums[n - 2] + nums[n - 1] < target {
			continue
		}
		for b := a + 1; b < n - 2; b++ {
			if b > a + 1 && nums[b] == nums[b - 1] {
				continue
			}
			if nums[a] + nums[b] + nums[b + 1] + nums[b + 2] > target {
				break
			}
			if nums[a] + nums[b] + nums[n - 2] + nums[n - 1] < target {
				continue
			}
			c, d := b + 1, n - 1
			for c < d {
				if nums[c] + nums[d] > target - (nums[a] + nums[b]) {
					d--
				} else if nums[c] + nums[d] < target - (nums[a] + nums[b]) {
					c++
				} else {
					res = append(res, []int{nums[a], nums[b], nums[c], nums[d]})
					for c < d && nums[c] == nums[c + 1] {
						c++
					}
					for c < d && nums[d] == nums[d - 1] {
						d--
					}
					c++
					d--
				}
			}
		}
	}

	return res
}

func main() {
	nums, target := []int{1,-2,-5,-4,-3,3,3,5}, -11
	res := fourSum(nums, target)
	fmt.Println(res)
}
