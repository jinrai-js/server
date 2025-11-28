package lru

import "errors"

type LRUCache struct {
	capacity   int
	cache      map[string]*node
	head, tail *node
}

type node struct {
	key, value string
	prev, next *node
}

func New(capacity int) LRUCache {
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
	if node, exists := lru.cache[key]; exists {
		lru.moveToHead(node)
		return node.value, nil
	}
	return "", errors.New("not found key: " + key)
}

func (lru *LRUCache) Put(key string, value string) {
	if node, exists := lru.cache[key]; exists {
		node.value = value
		lru.moveToHead(node)
		return
	}

	newNode := &node{key: key, value: value}
	lru.cache[key] = newNode
	lru.add(newNode)

	if len(lru.cache) > lru.capacity {
		tail := lru.popTail()
		delete(lru.cache, tail.key)
	}
}

func (lru *LRUCache) add(node *node) {
	node.prev = lru.head
	node.next = lru.head.next

	lru.head.next.prev = node
	lru.head.next = node
}

func (lru *LRUCache) remove(node *node) {
	prev := node.prev
	next := node.next

	prev.next = next
	next.prev = prev
}

func (lru *LRUCache) moveToHead(node *node) {
	lru.remove(node)
	lru.add(node)
}

func (lru *LRUCache) popTail() *node {
	node := lru.tail.prev
	lru.remove(node)
	return node
}
