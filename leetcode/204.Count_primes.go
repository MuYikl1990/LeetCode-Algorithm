package main

import "fmt"

func countPrimes(n int) int {
	count, notPrime := 0, make([]bool, n)

	for i := 2; i * i < n; i++ {
		if !notPrime[i] {
			for k := i * i; k < n; k += i {
				notPrime[k] = true
			}
		}
	}

	for i := 2; i < n; i++ {
		if !notPrime[i] {
			count++
		}
	}
	return count
}

func main() {
	n := 10
	res := countPrimes(n)
	fmt.Println(res)
}
