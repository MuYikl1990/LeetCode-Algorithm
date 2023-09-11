package main

import "fmt"

func subarraysWithKDistinct(nums []int, k int) int {
	var mostKSubarray func(int) int
	mostKSubarray = func(n int) int {
		count, nMap, res := 0, make(map[int]int), 0
		left, right := 0, 0
		for right < len(nums) {
			nMap[nums[right]]++
			if nMap[nums[right]] == 1 {
				count++
			}
			for count > n {
				nMap[nums[left]]--
				if nMap[nums[left]] == 0 {
					count--
				}
				left++
			}
			res += right - left + 1
			right++
		}
		return res
	}
	return mostKSubarray(k) - mostKSubarray(k - 1)
}

func main() {
	nums, k := []int{1,2,1,2,3}, 2
	res := subarraysWithKDistinct(nums, k)
	fmt.Println(res)
}
