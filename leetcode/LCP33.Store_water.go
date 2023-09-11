package main

import (
	"fmt"
	"math"
)

func storeWater(bucket []int, vat []int) int {
	res, right := math.MaxInt, 0
	for i := range vat {
		if vat[i] > right {
			right = vat[i]
		}
	}
	if right == 0 {
		return 0
	}

	for k := 1; k <= right; k++ {
		tmp := 0
		for i := 0; i < len(bucket); i++ {
			tmp += max33(0, (vat[i] + k - 1) / k - bucket[i])
		}
		res = min33(res, tmp + k)
	}
	return res
}

func min33(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max33(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	bucket, vat := []int{9,0,1}, []int{0,2,2}
	res := storeWater(bucket, vat)
	fmt.Println(res)
}
