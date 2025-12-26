package lru

import "errors"

type LRUCache struct {
	capacity   int
	cache      map[string]*node
	head, tail *node
}

type node struct {
	key        string
	value      string
	prev, next *node
}

func New(capacity int) LRUCache {
	if capacity <= 0 {
		capacity = 1
	}

	lru := LRUCache{
		capacity: capacity,
		cache:    make(map[string]*node),
		head:     &node{},
		tail:     &node{},
	}
	lru.head.next = lru.tail
	lru.tail.prev = lru.head
	return lru
}

func (lru *LRUCache) Get(key string) (string, error) {
	if node, exists := lru.cache[key]; exists && node != nil {
		lru.moveToHead(node)
		return node.value, nil
	}
	return "", errors.New("not found key: " + key)
}

func (lru *LRUCache) Put(key string, value string) {
	if lru.capacity <= 0 {
		return
	}

	if node, exists := lru.cache[key]; exists {
		node.value = value
		lru.moveToHead(node)
		return
	}

	newNode := &node{key: key, value: value}
	lru.cache[key] = newNode
	lru.add(newNode)

	if len(lru.cache) > lru.capacity {
		if tail := lru.popTail(); tail != nil {
			delete(lru.cache, tail.key)
		}
	}
}

func (lru *LRUCache) Has(key string) bool {
	_, exists := lru.cache[key]
	return exists
}

func (lru *LRUCache) add(node *node) {
	if node == nil || lru.head == nil || lru.head.next == nil {
		return
	}

	node.prev = lru.head
	node.next = lru.head.next

	lru.head.next.prev = node
	lru.head.next = node
}

func (lru *LRUCache) remove(node *node) {
	if node == nil || node.prev == nil || node.next == nil {
		return
	}

	prev := node.prev
	next := node.next

	prev.next = next
	next.prev = prev
}

func (lru *LRUCache) moveToHead(node *node) {
	if node == nil {
		return
	}
	lru.remove(node)
	lru.add(node)
}

func (lru *LRUCache) popTail() *node {
	if lru.tail == nil || lru.tail.prev == nil || lru.tail.prev == lru.head {
		return nil
	}

	node := lru.tail.prev
	lru.remove(node)
	return node
}
