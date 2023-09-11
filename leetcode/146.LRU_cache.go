package main

import "math"

type LRUCache struct {
	capacity int
	size int
	cache map[int]*DListNode
	head, tail *DListNode
}

type DListNode struct {
	key, value int
	prev, next *DListNode
}

func Constructor(capacity int) LRUCache {
	lru := LRUCache{
		capacity: capacity,
		cache: make(map[int]*DListNode),
		head: &DListNode{key: math.MinInt, value: 0},
		tail: &DListNode{key: math.MinInt, value: 0},
	}
	lru.head.next = lru.tail
	lru.tail.prev = lru.head
	return lru
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.cache[key]; !ok {
		return -1
	} else {
		this.MoveToHead(node)
		return node.value
	}
}

func (this *LRUCache) MoveToHead(node *DListNode) {
	this.RemoveNode(node)
	this.AddToHead(node)
}

func (this *LRUCache) RemoveNode(node *DListNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

func (this *LRUCache) Put(key, value int) {
	if node, ok := this.cache[key]; ok {
		node.value = value
		this.MoveToHead(node)
	} else {
		newNode := &DListNode{key: key, value: value}
		this.cache[key] = newNode
		this.AddToHead(newNode)
		this.size++
		if this.size > this.capacity {
			removed := this.RemoveTail()
			delete(this.cache, removed.key)
			this.size--
		}
	}
}

func (this *LRUCache) AddToHead(node *DListNode) {
	node.next = this.head.next
	node.prev = this.head
	this.head.next.prev = node
	this.head.next = node
}

func (this *LRUCache) RemoveTail() *DListNode {
	node := this.tail.prev
	this.RemoveNode(node)
	return node
}
