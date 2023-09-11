package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func atMostNGivenDigitSet(digits []string, n int) int {
	str := strings.Split(strconv.Itoa(n), "")
	digNum, nNum := len(digits), len(str)
	res := 0
	for i := 1; i < nNum; i++ {
		res += int(math.Pow(float64(digNum), float64(nNum - i)))
	}

	for i := 0; i < nNum; i++ {
		isCompare := false
		for _, digit := range digits {
			if digit < str[i] {
				res += int(math.Pow(float64(digNum), float64(nNum - i - 1)))
			} else if digit == str[i] {
				isCompare = true
				break
			}
		}
		if !isCompare {
			return res
		}
	}

	return res + 1
}

func main() {
	digits, n := []string{"1","4","9"}, 1000000000
	res := atMostNGivenDigitSet(digits, n)
	fmt.Println(res)
}
