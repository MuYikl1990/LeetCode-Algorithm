package main

import "fmt"

func numberOfSubarrays(nums []int, k int) int {
	left, right := 0, 0
	oddCnt, res := 0, 0

	for right < len(nums) {
		if nums[right] % 2 == 1 {
			oddCnt++
		}
		right++
		if oddCnt == k {
			index := right - 1
			for right < len(nums) && nums[right] % 2 == 0 {
				right++
			}
			rCnt := right - index
			lCnt := 1
			for nums[left] % 2 == 0 {
				lCnt++
				left++
			}
			res += lCnt * rCnt
			left++
			oddCnt--
		}
	}
	return res
}

func main() {
	nums, k := []int{2,2,2,1,2,2,1,2,2,2}, 2
	res := numberOfSubarrays(nums, k)
	fmt.Println(res)
}
