package main

import (
	"fmt"
	"strconv"
	"strings"
)

func ipToInt(ip string) int {
	strs := strings.Split(ip, ".")
	var nums [4]int
	for i, str := range strs {
		nums[i], _ = strconv.Atoi(str)
	}
	var res int
	for _, num := range nums {
		res = res << 8 | num
	}
	return res
}

func intToIp(num int) string {
	mod := 255
	nums := [4]int{}
	for i := 3; i >= 0; i-- {
		nums[i] = num & mod
		num >>= 8
	}
	var res []string
	for _, num := range nums {
		res = append(res, strconv.Itoa(num))
	}
	return strings.Join(res, ".")
}

func main() {
	ip := "10.0.3.193"
	num := ipToInt(ip)
	fmt.Println(num)
	ip = intToIp(num)
	fmt.Println(ip)
}
