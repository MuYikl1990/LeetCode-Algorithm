package main

import "fmt"

func maxSumTwoNoOverlap(nums []int, firstLen int, secondLen int) int {
	preSum := make([]int, len(nums) + 1)
	for i := 1; i < len(preSum); i++ {
		preSum[i] = preSum[i - 1] + nums[i - 1]
	}

	res, maxA, maxB := 0, 0, 0
	for i := firstLen + secondLen; i < len(preSum); i++ {
		maxA = max1031(maxA, preSum[i - secondLen] - preSum[i - secondLen - firstLen])
		maxB = max1031(maxB, preSum[i - firstLen] - preSum[i - secondLen - firstLen])
		res = max1031(res, max1031(maxA + preSum[i] - preSum[i - secondLen], maxB + preSum[i] - preSum[i - firstLen]))
	}
	return res
}

func max1031(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	nums, firstLen, secondLen := []int{2,1,5,6,0,9,5,0,3,8}, 4, 3
	res := maxSumTwoNoOverlap(nums, firstLen, secondLen)
	fmt.Println(res)
}
