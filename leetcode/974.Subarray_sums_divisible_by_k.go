package main

import "fmt"

// 前缀和
func subarraysDivByK(nums []int, k int) int {
	res, arrMod := 0, 0
	arrMap := make(map[int]int)
	arrMap[0] = 1

	for _, num := range nums {
		arrMod = (arrMod + num) % k
		if arrMod < 0 {
			arrMod += k
		}
		if val, ok := arrMap[arrMod]; ok {
			res += val
		}
		arrMap[arrMod]++
	}
	return res
}

func main() {
	nums, k := []int{4,5,0,-2,-3,1}, 5
	res := subarraysDivByK(nums, k)
	fmt.Println(res)
}
