package main

import "fmt"

type LFUCache struct {
	capacity int
	size int
	kvCache map[int]*Node460
	freqCache map[int]*DList
	minFreq int
}

type Node460 struct {
	key, value, freq int
	prev, next *Node460
}

type DList struct {
	head, tail *Node460
}

func initNode460(key, value, freq int) *Node460 {
	return &Node460{key: key, value: value, freq: freq}
}

func initDList() *DList {
	dlist := &DList{
		head: &Node460{key: -1, value: -1, freq: -1},
		tail: &Node460{key: -1, value: -1, freq: -1},
	}
	dlist.head.next = dlist.tail
	dlist.tail.prev = dlist.head
	return dlist
}

func (dlist *DList) AddToHead(Node460 *Node460) {
	Node460.prev = dlist.head
	Node460.next = dlist.head.next
	dlist.head.next.prev = Node460
	dlist.head.next = Node460
}

func (dlist *DList) RemoveNode460(Node460 *Node460) {
	Node460.prev.next = Node460.next
	Node460.next.prev = Node460.prev
	Node460.prev, Node460.next = nil, nil
}

func (dlist *DList) RemoveTail() *Node460 {
	Node460 := dlist.tail.prev
	dlist.tail.prev = Node460.prev
	Node460.prev.next = dlist.tail
	return Node460
}

func (dlist *DList) isEmpty() bool {
	return dlist.head.next == dlist.tail
}


func Constructor460(capacity int) LFUCache {
	lfu := LFUCache {
		capacity: capacity,
		kvCache: make(map[int]*Node460),
		freqCache: make(map[int]*DList),
	}

	return lfu
}


func (this *LFUCache) Get(key int) int {
	if Node460, ok := this.kvCache[key]; ok {
		this.increment(Node460, false)
		return Node460.value
	}
	return -1
}


func (this *LFUCache) Put(key int, value int)  {
	if Node460, ok := this.kvCache[key]; ok {
		Node460.value = value
		this.increment(Node460, false)
	} else {
		if this.capacity == 0 {
			return
		}
		if this.size == this.capacity {
			this.removeMinFreqNode460()
		}
		Node460 := initNode460(key, value, 1)
		this.increment(Node460, true)
		this.kvCache[key] = Node460
	}
}

func (this *LFUCache) increment(Node460 *Node460, isNewNode460 bool) {
	if isNewNode460 {
		this.minFreq = 1
		this.addToDList(Node460)
		this.size++
	} else {
		this.removeNode460(Node460)
		Node460.freq++
		this.addToDList(Node460)
		if _, ok := this.freqCache[this.minFreq]; !ok {
			this.minFreq++
		}
	}
}

func (this *LFUCache) addToDList(Node460 *Node460) {
	if _, ok := this.freqCache[Node460.freq]; !ok {
		this.freqCache[Node460.freq] = initDList()
	}
	this.freqCache[Node460.freq].AddToHead(Node460)
}

func (this *LFUCache) removeNode460(Node460 *Node460) {
	list := this.freqCache[Node460.freq]
	list.RemoveNode460(Node460)
	if list.isEmpty() {
		delete(this.freqCache, Node460.freq)
	}
}

func (this *LFUCache) removeMinFreqNode460() {
	list := this.freqCache[this.minFreq]
	last := list.RemoveTail()
	delete(this.kvCache, last.key)
	if list.isEmpty() {
		delete(this.freqCache, this.minFreq)
	}
	this.size--
}

func main() {
	capacity, key, value := 2, 3, 2
	obj := Constructor460(capacity)
	param := obj.Get(key)
	fmt.Println(param)
	obj.Put(key,value)
}
