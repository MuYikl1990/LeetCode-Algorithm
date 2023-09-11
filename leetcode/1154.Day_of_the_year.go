package main

import (
	"fmt"
	"strconv"
)

func dayOfYear(date string) int {
	year, _ := strconv.Atoi(date[:4])
	month, _ := strconv.Atoi(date[5:7])
	day, _ := strconv.Atoi(date[8:])

	days := [12]int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	if year % 400 == 0 || (year % 4 == 0 && year % 100 != 0) {
		days[1]++
	}

	res := 0
	for i := 0; i < month - 1; i++ {
		res += days[i]
	}
	return res + day
}

func main() {
	date := "2019-03-09"
	res := dayOfYear(date)
	fmt.Println(res)
}
