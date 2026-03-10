package main

type DoubleNode struct {
	key, val int
	prev     *DoubleNode
	next     *DoubleNode
}

type LRUCache struct {
	cap   int
	cache map[int]*DoubleNode
	left  *DoubleNode
	right *DoubleNode
}

func Constructor(capacity int) LRUCache {
	left := &DoubleNode{0, 0, nil, nil}
	right := &DoubleNode{0, 0, nil, nil}

	left.next = right
	right.prev = left

	return LRUCache{
		cap:   capacity,
		cache: make(map[int]*DoubleNode),
		left:  left,
		right: right,
	}
}

func (this *LRUCache) Get(key int) int {
	node, ok := this.cache[key]
	if ok {
		this.remove(node)
		this.insert(node)
		return node.val
	}

	return -1
}

func (this *LRUCache) Put(key int, value int) {
	node, ok := this.cache[key]
	if ok {
		this.remove(node)
	}

	node = &DoubleNode{key, value, nil, nil}
	this.cache[key] = node
	this.insert(node)

	if len(this.cache) > this.cap {
		lru := this.left.next
		this.remove(lru)
		delete(this.cache, lru.key)
	}
}

func (this *LRUCache) remove(node *DoubleNode) {
	prev, next := node.prev, node.next
	prev.next, next.prev = next, prev
}

func (this *LRUCache) insert(node *DoubleNode) {
	prev, next := this.right.prev, this.right
	prev.next, next.prev = node, node
	node.prev, node.next = prev, next
}
