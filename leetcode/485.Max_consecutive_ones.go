package main

import "fmt"

func findMaxConsecutiveOnes(nums []int) int {
	res, tmp := 0, 0
	for i := range nums {
		if nums[i] == 1 {
			tmp++
		} else {
			res = max485(res, tmp)
			tmp = 0
		}
	}
	return max485(res, tmp)
}

func max485(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	nums := []int{1,0,1,1,0,1}
	res := findMaxConsecutiveOnes(nums)
	fmt.Println(res)
}
