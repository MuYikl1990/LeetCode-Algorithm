package main

import (
	"fmt"
	"math"
)

type TreeAncestor struct {
	dp  [][]int
	dep int
}

// dp倍增算法
func Constructor1483(n int, parent []int) TreeAncestor {
	dep := int(math.Log(float64(n)) / math.Log(2)) + 1
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, dep + 1)
	}
	for i := 0; i < n; i++ {
		for j := 0; j <= dep; j++ {
			dp[i][j] = -1
			if j == 0 {
				dp[i][j] = parent[i]
			}
		}
	}

	for i := 0; i < n; i++ {
		for j := 1; j <= dep; j++ {
			if dp[i][j - 1] != -1 {
				dp[i][j] = dp[dp[i][j - 1]][j - 1]
			}
		}
	}

	return TreeAncestor{dp, dep}
}


func (this *TreeAncestor) GetKthAncestor(node int, k int) int {
	for i := this.dep; i >= 0; i-- {
		if k & (1 << i) != 0 {
			node = this.dp[node][i]
			if node == -1 {
				break
			}
		}
	}
	return node
}

func main() {
	n, parent := 7, []int{-1, 0, 0, 1, 1, 2, 2}
	obj := Constructor1483(n, parent)
	node, k := 6, 3
	param := obj.GetKthAncestor(node,k)
	fmt.Println(param)
}
