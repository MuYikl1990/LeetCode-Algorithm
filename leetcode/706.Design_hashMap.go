package main

import "fmt"

const base = 997

type entry struct {
	next     *entry
	key, val int
}

type MyHashMap struct {
	data [base]*entry
}


func Constructor706() MyHashMap {
	return MyHashMap{}
}


func (this *MyHashMap) Put(key int, value int)  {
	index := key % base
	node := this.data[index]
	if node == nil {
		this.data[index] = &entry{key: key, val: value}
		return
	}

	for {
		if node.key == key {
			node.val = value
			return
		}

		if node.next == nil {
			node.next = &entry{key: key, val: value}
			return
		}
		node = node.next
	}
}


func (this *MyHashMap) Get(key int) int {
	index := key % base
	node := this.data[index]

	for node != nil {
		if node.key == key {
			return node.val
		}
		node = node.next
	}
	return -1
}


func (this *MyHashMap) Remove(key int)  {
	index := key % base
	node := this.data[index]

	if node == nil {
		return
	}

	if node.key == key {
		this.data[index] = this.data[index].next
		return
	}

	pre := node
	node = node.next
	for node != nil {
		if node.key == key {
			pre.next = node.next
			return
		}
		pre = node
		node = node.next
	}
}

func main() {
	key1, value1 := 2, 9
	key2, value2 := 9107, 67880
	obj := Constructor706()
	obj.Put(key1,value1)
	param1 := obj.Get(key1)
	fmt.Println(param1)
	obj.Put(key2,value2)
	param2 := obj.Get(key2)
	fmt.Println(param2)
	obj.Remove(key1)
	fmt.Println(obj.data[2])
}
