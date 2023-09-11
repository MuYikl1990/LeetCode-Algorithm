package main

import (
	"fmt"
	"math/bits"
)

func minNumberOfSemesters(n int, relations [][]int, k int) int {
	// 计算每门课的先修课集合
	before := make([]int, n)
	for _, d := range relations {
		before[d[1]-1] |= 1 << (d[0] - 1)
	}
	m := 1 << n
	// 计算所有课程集合的先修课集合，不合法的标记为 -1
	totPre := make([]int, m)
	for i := range totPre {
		if bits.OnesCount(uint(i)) > k {
			totPre[i] = -1
			continue
		}
		for s := uint(i); s > 0; s &= s - 1 {
			p := before[bits.TrailingZeros(s)]
			if p&i > 0 {
				totPre[i] = -1
				break
			}
			totPre[i] |= p
		}
	}
	dp := make([]int, 1<<n)
	for i := range dp {
		dp[i] = n
	}
	dp[0] = 0
	for s, v := range dp {
		t := m - 1 ^ s // 补集
		for sub := t; sub > 0; sub = (sub - 1) & t { // 枚举下个学期要学的课
			if p := totPre[sub]; p >= 0 && s&p == p { // 这些课的先修课必须合法且在之前的学期里必须上过
				dp[s|sub] = min1494(dp[s|sub], v+1)
			}
		}
	}
	return dp[m-1]
}

func min1494(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	n, relations, k := 5, [][]int{{2, 1}, {3, 1}, {4, 1}, {1, 5}}, 2
	res := minNumberOfSemesters(n, relations, k)
	fmt.Println(res)
}
