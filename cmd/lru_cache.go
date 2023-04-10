package main

type LRUCache struct {
	content map[int]*LRUNode
	head    *LRUNode
	tail    *LRUNode
	len     int
}

type LRUNode struct {
	left  *LRUNode
	right *LRUNode
	value int
	key   int
}

func Constructor(capacity int) LRUCache {
	l := LRUCache{}
	l.content = make(map[int]*LRUNode, capacity)
	l.len = capacity
	return l
}

//lint:ignore ST1006 this
func (this *LRUCache) Get(key int) int {
	if this.content[key] == nil {
		return -1
	}

	if this.head != this.content[key] {
		this.content[key].left.right = this.content[key].right

		if this.content[key].right != nil {
			this.content[key].right.left = this.content[key].left
		} else {
			this.tail = this.content[key].left
		}

		this.head.left = this.content[key]
		this.content[key].left = nil
		this.content[key].right = this.head
		this.head = this.content[key]
	}

	return this.content[key].value
}

//lint:ignore ST1006 this
func (this *LRUCache) Put(key int, value int) {
	if this.content[key] != nil {
		this.content[key].value = value
		this.Get(key)
	} else {
		if len(this.content) == this.len {
			delete(this.content, this.tail.key)
			if this.tail.left != nil {
				this.tail.left.right = nil
			}
			this.tail = this.tail.left
		}
		l := &LRUNode{}
		l.value = value
		l.key = key
		l.left = nil
		l.right = this.head
		if this.head != nil {
			this.head.left = l
		}
		this.head = l
		if this.tail == nil {
			this.tail = l
		}
		this.content[key] = l
	}
}
