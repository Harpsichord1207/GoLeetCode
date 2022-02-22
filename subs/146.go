package subs

import (
	"fmt"
)

type Node146 struct {
	value int
	next  *Node146
}

type LRUCache struct {
	data     map[int]int
	size     int
	capacity int
	head     *Node146
}

func popLastNode(node *Node146) int {
	value := -1
	if node.next == nil {
		return value
	}

	a := &Node146{next: node}
	b := node

	for {
		if b.next == nil {
			a.next = nil
			value = b.value
			break
		}
		a = b
		b = b.next
	}
	return value
}

func printLink(node *Node146) {
	var values []int
	n := node
	c := 0
	for {
		if n == nil {
			break
		}
		c += 1
		if c >= 10000 {
			panic("STOPSTOP")
		}
		values = append(values, n.value)
		n = n.next
	}
	fmt.Printf("Link: %d\n", values)

}

func Constructor146(capacity int) LRUCache {
	lruCache := LRUCache{}
	lruCache.capacity = capacity
	lruCache.data = make(map[int]int)
	return lruCache
}

func (this *LRUCache) Get(key int) int {
	value, ok := this.data[key]
	if !ok {
		return -1
	}

	if this.size <= 1 {
		return value
	}

	if this.head.value == key {
		return value
	}

	a := &Node146{next: this.head}
	b := this.head
	for {
		// fmt.Printf("In get iter: a = %d, b = %d, key = %d\n", a.value, b.value, key)
		if b.value == key {
			a.next = b.next
			b.next = this.head
			this.head = b
			break
		}
		a = b
		b = b.next
	}
	return value
}

func (this *LRUCache) Put(key int, value int) {

	_, ok := this.data[key]
	this.data[key] = value
	if ok {
		this.Get(key)
		return
	}

	newNode := &Node146{value: key}
	if this.head != nil {
		newNode.next = this.head
	}
	this.head = newNode

	this.size += 1

	if this.size > this.capacity {
		delete(this.data, popLastNode(this.head))
	}

}

func Test146() {

	// lruCache := Constructor146(2)
	// fmt.Println(lruCache.Get(2))
	// lruCache.Put(2, 6)
	// printLink(lruCache.head)
	// fmt.Println(lruCache.Get(1))
	// lruCache.Put(1, 5)
	// printLink(lruCache.head)
	// lruCache.Put(1, 2)
	// printLink(lruCache.head)
	// fmt.Println(lruCache.Get(1))
	// fmt.Println(lruCache.Get(2))

	lruCache := Constructor146(2)
	lruCache.Put(2, 1)
	printLink(lruCache.head)
	lruCache.Put(1, 1)
	printLink(lruCache.head)
	lruCache.Put(2, 3)
	printLink(lruCache.head)
	lruCache.Put(4, 1)
	printLink(lruCache.head)
	fmt.Println(lruCache.Get(1))
	fmt.Println(lruCache.Get(2))
}
