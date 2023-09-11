package main

import (
	"fmt"
)

func compareVersion(version1 string, version2 string) int {
	m, n := len(version1), len(version2)
	i, j := 0, 0
	for i < m || j < n {
		x := 0
		for i < m && version1[i] != '.' {
			x = x * 10 + int(version1[i] - '0')
			i++
		}
		i++
		y := 0
		for j < n && version2[j] != '.' {
			y = y * 10 + int(version2[j] - '0')
			j++
		}
		j++
		if x != y {
			if x > y {
				return 1
			} else {
				return -1
			}
		}
	}
	return 0
}

func main() {
	version1, version2 := "1.011", "1.001"
	res := compareVersion(version1, version2)
	fmt.Println(res)
}
