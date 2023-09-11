package main

import (
	"fmt"
	"math"
)

func findMedianSortedArrays(nums1, nums2 []int) float64 {
	n := len(nums1)
	m := len(nums2)

	if n > m {
		return findMedianSortedArrays(nums2, nums1)
	}

	lMax1, rMin1, lMax2, rMin2 := 0, 0, 0, 0
	c1, c2, lo, hi := 0, 0, 0, n

	for lo <= hi {
		c1 = (lo + hi) / 2
		c2 = (n + m + 1) / 2 - c1

		if c1 == 0 {
			lMax1 = math.MinInt
		} else {
			lMax1 = nums1[c1 - 1]
		}

		if c1 == n {
			rMin1 = math.MaxInt
		} else {
			rMin1 = nums1[c1]
		}

		if c2 == 0 {
			lMax2 = math.MinInt
		} else {
			lMax2 =nums2[c2 - 1]
		}

		if c2 == m {
			rMin2 = math.MaxInt
		} else {
			rMin2 = nums2[c2]
		}

		if lMax1 > rMin2 {
			hi = c1 - 1
		} else if lMax2 > rMin1 {
			lo = c1 + 1
		} else {
			break
		}
	}

	if (m + n) % 2 == 1 {
		return Maxf(lMax1, lMax2)
	}
	return (Maxf(lMax1, lMax2) + Minf(rMin1, rMin2)) / 2
}

func Maxf(a, b int) float64 {
	if a > b {
		return float64(a)
	}
	return float64(b)
}

func Minf(a, b int) float64 {
	if a < b {
		return float64(a)
	}
	return float64(b)
}

func main() {
	nums1 := []int{1, 3}
	nums2 := []int{2}
	res := findMedianSortedArrays(nums1, nums2)
	fmt.Println(res)
}
