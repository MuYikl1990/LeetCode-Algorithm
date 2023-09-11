package main

import (
	"fmt"
	"math/rand"
)

type Solution struct {
	nums []int
}


func Constructor384(nums []int) Solution {
	return Solution{nums}
}


func (this *Solution) Reset() []int {
	return this.nums
}


func (this *Solution) Shuffle() []int {
	temp := make([]int, len(this.nums))
	copy(temp, this.nums)
	for i := 0; i < len(temp); i++ {
		idx := rand.Intn(len(temp) - i) + i
		temp[i], temp[idx] = temp[idx], temp[i]
	}
	return temp
}

func main() {
	nums := []int{1,2,3,4,5,6,7}
	obj := Constructor384(nums)
	param1 := obj.Reset()
	fmt.Println(param1)
	param2 := obj.Shuffle()
	fmt.Println(param2)
	param3 := obj.Shuffle()
	fmt.Println(param3)
}
