package lrucache

type LRUCache struct {
	hash       map[int]*Node
	head, tail *Node
	capacity   int
}

type Node struct {
	prev, next *Node
	key, val   int
}

func Constructor(capacity int) LRUCache {
	cache := LRUCache{
		hash:     map[int]*Node{},
		head:     &Node{},
		tail:     &Node{},
		capacity: capacity,
	}

	cache.head.next = cache.tail
	cache.tail.prev = cache.head

	return cache
}

func (this *LRUCache) Get(key int) int {
	if value, ok := this.hash[key]; ok {
		value.next.prev = value.prev
		value.prev.next = value.next
		value.next = this.head.next
		value.prev = this.head
		this.head.next.prev = value
		this.head.next = value

		return value.val
	}

	return -1
}

func (this *LRUCache) Put(key int, value int) {
	var (
		node *Node
		ok   bool
	)
	if node, ok = this.hash[key]; ok {
		node.val = value
		node.prev.next = node.next
		node.next.prev = node.prev
	} else {
		node = &Node{
			key: key,
			val: value,
		}
	}

	// 将 node 插入首部
	node.next = this.head.next
	node.prev = this.head
	this.head.next.prev = node
	this.head.next = node
	this.hash[key] = node

	if len(this.hash) > this.capacity {
		//超了，需要将最后一个元素删除
		delNode := this.tail.prev
		this.tail.prev = delNode.prev
		delNode.prev.next = this.tail
		delete(this.hash, delNode.key)
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
