package main

import (
	"fmt"
	"math/rand"
)

type RandomizedSet struct {
	numMap map[int]int
	arr    []int
	idx    int
}


func Constructor380() RandomizedSet {
	return RandomizedSet{make(map[int]int), []int{}, -1}
}


func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.numMap[val]; ok {
		return false
	} else {
		this.arr = append(this.arr, val)
		this.idx++
		this.numMap[val] = this.idx
		return true
	}
}


func (this *RandomizedSet) Remove(val int) bool {
	if pos, ok := this.numMap[val]; !ok {
		return false
	} else {
		if pos != this.idx {
			this.arr[pos], this.arr[this.idx] = this.arr[this.idx], this.arr[pos]
			this.numMap[this.arr[pos]] = pos
		}
		delete(this.numMap, this.arr[this.idx])
		this.arr = this.arr[:len(this.arr) - 1]
		this.idx--
		return true
	}
}


func (this *RandomizedSet) GetRandom() int {
	pos := rand.Intn(this.idx + 1)
	return this.arr[pos]
}


func main() {
	obj := Constructor380()
	val1, val2 := 1, 2
	param := obj.Insert(val1)
	fmt.Println(param)
	param = obj.Remove(val2)
	fmt.Println(param)
	param = obj.Insert(val2)
	fmt.Println(param)
	param = obj.Remove(val2)
	fmt.Println(param)
	res := obj.GetRandom()
	fmt.Println(res)
}
