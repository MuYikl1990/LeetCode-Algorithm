package main

import "fmt"

func combine(n int, k int) [][]int {
	var res [][]int
	backtrace77(n, k, 1, []int{}, &res)
	return res
}

func backtrace77(n, k, index int, tmp []int, res *[][]int) {
	if len(tmp) == k {
		com := make([]int, len(tmp))
		copy(com, tmp)
		*res = append(*res, com)
		return
	}

	if index > n || len(tmp) + n - index < k - 1 {
		return
	}
	for i := index; i <= n; i++ {
		tmp = append(tmp, i)
		backtrace77(n, k, i + 1, tmp, res)
		tmp = tmp[:len(tmp) - 1]
	}

}

func main() {
	n, k := 4, 2
	res := combine(n, k)
	fmt.Println(res)
}
