package main

import (
	"fmt"
	"math"
	"strconv"
)

func nextGreaterElement(n int) int {
	str := []byte(strconv.Itoa(n))
	if len(str) < 2 {
		return -1
	}
	i, left, right := len(str) - 2, len(str) - 1, len(str) - 1
	for i >= 0 && str[i] >= str[left] {
		i--
		left--
	}

	if i >= 0 {
		for str[i] >= str[right] {
			right--
		}
		str[i], str[right] = str[right], str[i]

		for start, end := left, len(str) - 1; start < end; start, end = start + 1, end - 1 {
			str[start], str[end] = str[end], str[start]
		}

		num, _ := strconv.Atoi(string(str))
		if float64(num) > math.Pow(2, 31) - 1 {
			return -1
		}
		return num
	}

	return -1

	//str := strings.Split(strconv.Itoa(n), "")
	//for i := len(str) - 1; i > 0; i-- {
	//	if str[i] >= str[i - 1] {
	//		sort.Strings(str[i:])
	//		for j := i; j < len(str); j++ {
	//			if str[j] >= str[i - 1] {
	//				str[j], str[i - 1] = str[i - 1], str[j]
	//				num, _ := strconv.Atoi(strings.Join(str, ""))
	//				if float64(num) > math.Pow(2, 31) - 1 {
	//					return -1
	//				}
	//				return num
	//			}
	//		}
	//	}
	//}
	//return -1
}

func main() {
	n := 12443322
	res := nextGreaterElement(n)
	fmt.Println(res)
}
