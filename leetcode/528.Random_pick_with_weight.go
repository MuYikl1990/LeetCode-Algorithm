package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Solution528 struct {
	maxInt int
	preSum []int
}


func Constructor528(w []int) Solution528 {
	preSum := make([]int, len(w))
	preSum[0] = w[0]
	for i := 1; i < len(w); i++ {
		preSum[i] = preSum[i - 1] + w[i]
	}
	return Solution528{maxInt: preSum[len(preSum) - 1], preSum: preSum}
}


func (this *Solution528) PickIndex() int {
	var binarySearch func(int) int
	binarySearch = func(target int) int {
		left, right := 0, len(this.preSum) - 1
		for left < right {
			mid := left + (right - left) / 2
			if this.preSum[mid] > target {
				right = mid
			} else {
				left = mid + 1
			}
		}
		return right
	}

	rand.Seed(time.Now().UnixNano())
	target := rand.Intn(this.maxInt)
	res := binarySearch(target)
	return res
}

func main() {
	w := []int{4,1,7}
	obj := Constructor528(w)
	param := obj.PickIndex()
	fmt.Println(param)
}
