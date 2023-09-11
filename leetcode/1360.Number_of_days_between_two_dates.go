package main

import (
	"fmt"
	"strconv"
	"strings"
)

func daysBetweenDates(date1 string, date2 string) int {
	return abs1360(daysTo(date1) - daysTo(date2))
}

func abs1360(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func daysTo(date string) int {
	res := 0
	days := []int{0, 31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	tmp := strings.Split(date, "-")
	year, _ := strconv.Atoi(tmp[0])
	month, _ := strconv.Atoi(tmp[1])
	day, _ := strconv.Atoi(tmp[2])

	res += day
	for month != 0 {
		month--
		res += days[month]
		if month == 2 && isLeapYear(year) {
			res++
		}
	}
	res += 365 * (year - 1971)
	res += (year - 1) / 400 - 1971 / 400
	res -= (year - 1) / 100 - 1971 / 100
	res += (year - 1) / 4 - 1971 / 4
	return res
}

func isLeapYear(year int) bool {
	return year % 400 == 0 || (year % 100 != 0 && year % 4 == 0)
}

func main() {
	date1, date2 := "2020-01-15", "2019-12-31"
	res := daysBetweenDates(date1, date2)
	fmt.Println(res)
}
