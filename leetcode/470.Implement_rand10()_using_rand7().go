package main

import (
	"fmt"
	"math/rand"
	"time"
)

func rand7() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(7) + 1
}

func rand10() int {
	res := 0
	for true {
		num := (rand7() - 1) * 7 + rand7()
		if num <= 40 {
			res = num % 10 + 1
			break
		}
	}
	return res
}

func main() {
	n := 3
	for n > 0 {
		fmt.Println(rand10())
		n -= 1
	}
}
