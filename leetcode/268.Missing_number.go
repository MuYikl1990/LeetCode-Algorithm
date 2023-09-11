package main

import "fmt"

func missingNumberI(nums []int) int {
	//res := 0
	//for i := 1; i <= len(nums); i++ {
	//	res ^= i
	//}
	//for i := 0; i < len(nums); i++ {
	//	res ^= nums[i]
	//}
	//return res
	res := 0
	for i := 0; i < len(nums); i++ {
		res += i - nums[i]
	}
	return res + len(nums)
}

func main() {
	nums := []int{9,6,4,2,3,5,7,0,1}
	res := missingNumberI(nums)
	fmt.Println(res)
}
