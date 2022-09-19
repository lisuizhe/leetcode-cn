package leetcode

/*
 * @lc app=leetcode.cn id=460 lang=golang
 *
 * [460] LFU 缓存
 */

// @lc code=start
type LFUCache struct {
	Keys              map[int]*Node
	Freqs             map[int]*DoubleList
	Capacity, MinFreq int
}

type Node struct {
	Key, Val, Freq int
	Prev, Next     *Node
}

type DoubleList struct {
	dummyHead, dummyTail *Node
}

func DLConstructor() *DoubleList {
	h, t := &Node{}, &Node{}
	h.Next = t
	t.Prev = h
	return &DoubleList{
		dummyHead: h,
		dummyTail: t,
	}
}

func (this *DoubleList) Add(node *Node) {
	node.Next = this.dummyHead.Next
	node.Prev = this.dummyHead
	this.dummyHead.Next.Prev = node
	this.dummyHead.Next = node
}

func (this *DoubleList) Remove(node *Node) {
	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev
	node.Prev = nil
	node.Next = nil
}

func (this *DoubleList) IsEmpty() bool {
	return this.dummyHead.Next == this.dummyTail
}

func (this *DoubleList) RemoveLast() *Node {
	if this.IsEmpty() {
		return nil
	}
	last := this.dummyTail.Prev
	this.Remove(last)
	return last
}

func Constructor(capacity int) LFUCache {
	return LFUCache{
		Keys:     make(map[int]*Node, capacity),
		Freqs:    make(map[int]*DoubleList),
		Capacity: capacity,
	}
}

func (this *LFUCache) Get(key int) int {
	if node, ok := this.Keys[key]; ok {
		this.incFreq(node)
		return node.Val
	}
	return -1
}

func (this *LFUCache) Put(key int, value int) {
	if this.Capacity <= 0 {
		return
	}

	if node, ok := this.Keys[key]; ok {
		node.Val = value
		this.incFreq(node)
	} else {
		// evict least frequantly used if over capacity
		if len(this.Keys) >= this.Capacity {
			evicted := this.Freqs[this.MinFreq].RemoveLast()
			delete(this.Keys, evicted.Key)
		}

		node := &Node{Key: key, Val: value, Freq: 1}
		this.Keys[key] = node
		if this.Freqs[1] == nil {
			this.Freqs[1] = DLConstructor()
		}
		this.Freqs[1].Add(node)
		this.MinFreq = 1
	}
}

func (this *LFUCache) incFreq(node *Node) {
	oldFreq := node.Freq
	// remove from Freqs
	this.Freqs[oldFreq].Remove(node)
	// update MinFreq if required
	if this.MinFreq == oldFreq && this.Freqs[oldFreq].IsEmpty() {
		this.MinFreq++
		delete(this.Freqs, oldFreq)
	}
	node.Freq++
	// add to Freqs
	if this.Freqs[node.Freq] == nil {
		this.Freqs[node.Freq] = DLConstructor()
	}
	this.Freqs[node.Freq].Add(node)
}

/**
 * Your LFUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
// @lc code=end
