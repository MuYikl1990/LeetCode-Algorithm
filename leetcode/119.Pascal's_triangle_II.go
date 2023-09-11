package main

import "fmt"

func getRow(rowIndex int) []int {
	var res []int
	res = append(res, 1)
	for i := 1; i <= rowIndex; i++ {
		for j := i - 1; j > 0; j-- {
			res[j] = res[j - 1] + res[j]
		}
		res = append(res, 1)
	}
	return res
}

func main() {
	rowIndex := 5
	res := getRow(rowIndex)
	fmt.Println(res)
}
