package main

import "fmt"

func findMaxLength(nums []int) int {
	res, sum := 0, 0
	numMap := make(map[int]int)
	numMap[0] = -1
	for i, num := range nums {
		if num == 0 {
			sum += -1
		} else {
			sum += 1
		}

		if pos, ok := numMap[sum]; ok {
			res = max525(res, i - pos)
		} else {
			numMap[sum] = i
		}
	}
	return res
}

func max525(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	nums := []int{0,0,1,0,0,0,1,1}
	res := findMaxLength(nums)
	fmt.Println(res)
}
