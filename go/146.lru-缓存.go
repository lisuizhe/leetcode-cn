package leetcode

/*
 * @lc app=leetcode.cn id=146 lang=golang
 *
 * [146] LRU 缓存
 */

// @lc code=start
type LRUCache struct {
	head, tail *Node
	Keys       map[int]*Node
	Capacity   int
}

type Node struct {
	Key, Val   int
	Prev, Next *Node
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		Keys:     make(map[int]*Node, capacity),
		Capacity: capacity,
	}
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.Keys[key]; ok {
		// Keys will not change, update nodes only
		this.removeNode(node)
		this.addNode(node)
		return node.Val
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.Keys[key]; ok {
		// update value of Keys[key]
		node.Val = value
		// update node
		this.removeNode(node)
		this.addNode(node)
	} else {
		node = &Node{Key: key, Val: value}
		// update Keys
		this.Keys[key] = node
		// update node
		this.addNode(node)
	}
	// check capacity
	if len(this.Keys) > this.Capacity {
		// update Keys
		delete(this.Keys, this.tail.Key)
		// update node
		this.removeNode(this.tail)
	}
}

func (this *LRUCache) addNode(node *Node) {
	// update node's Prev/Next
	node.Prev = nil
	node.Next = this.head
	// set current head's Prev to node if head is not nil
	if this.head != nil {
		this.head.Prev = node
	}
	// set head to node
	this.head = node
	// set tail if current tail is nil
	if this.tail == nil {
		this.tail = node
	}
}

func (this *LRUCache) removeNode(node *Node) {
	if node == this.head && node == this.tail {
		this.head, this.tail = nil, nil
	} else if node == this.head {
		this.head = node.Next
		this.head.Prev = nil
	} else if node == this.tail {
		this.tail = node.Prev
		this.tail.Next = nil
	} else {
		node.Prev.Next = node.Next
		node.Next.Prev = node.Prev
	}
	// clean up; decrease memory but increase runtime
	// node.Prev = nil
	// node.Next = nil
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
// @lc code=end
