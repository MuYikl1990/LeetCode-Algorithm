package main

import "fmt"

func findKthNumber(n int, k int) int {
	var getCount func(int, int) int
	getCount = func(prefix, n int) int {
		count := 0
		for cur, next := prefix, prefix + 1; cur <= n; cur, next = cur * 10, next * 10 {
			count += min440(n + 1, next) - cur
		}
		return count
	}
	p, prefix := 1, 1  // p作为一个指针，指向当前所在位置，当p==k时，也就是到了排位第k的数
	for p < k {
		count := getCount(prefix, n)  // 获得当前前缀下所有子节点的和
		if p + count > k {  // 第k个数在当前前缀下
			prefix *= 10
			p++  // 把指针指向了第一个子节点的位置，比如11乘10后变成110，指针从11指向了110
		} else {  // 第k个数不在当前前缀下
			prefix++
			p += count  // 注意这里的操作，把指针指向了下一前缀的起点
		}
	}
	return prefix
}

func min440(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	n, k := 27, 15
	res := findKthNumber(n, k)
	fmt.Println(res)
}
