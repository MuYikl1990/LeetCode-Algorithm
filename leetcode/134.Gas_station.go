package main

import (
	"fmt"
	"math"
)

func canCompleteCircuit(gas []int, cost []int) int {
	remain := 0
	minGas := math.MaxInt
	res := 0
	for i := 0; i < len(gas); i++ {
		remain += gas[i] - cost[i]
		if remain < minGas {
			minGas = remain
			res = i
		}
	}
	if remain < 0 {
		return -1
	}
	return (res + 1) % len(gas)
}

func main() {
	gas, cost := []int{1,2,3,4,5}, []int{3,4,5,1,2}
	res := canCompleteCircuit(gas, cost)
	fmt.Println(res)
}
