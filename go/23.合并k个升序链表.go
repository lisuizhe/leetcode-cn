package leetcode

import "container/heap"

/*
 * @lc app=leetcode.cn id=23 lang=golang
 *
 * [23] 合并K个升序链表
 */

// @lc code=start
//
// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	return mergeKListsByRecursion(lists)
	//return mergeKListsByHeap(lists)
}

type IntHeap []int

func (h IntHeap) Len() int            { return len(h) }
func (h IntHeap) Less(i, j int) bool  { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func mergeKListsByHeap(lists []*ListNode) *ListNode {
	h := &IntHeap{}
	for _, l := range lists {
		p := l
		for p != nil {
			heap.Push(h, p.Val)
			p = p.Next
		}
	}
	dummy := &ListNode{Val: -1, Next: nil}
	p := dummy
	for len(*h) > 0 {
		x := heap.Pop(h)
		p.Next = &ListNode{Val: x.(int), Next: nil}
		p = p.Next
	}
	return dummy.Next
}

func mergeKListsByRecursion(lists []*ListNode) *ListNode {
	n := len(lists)
	if n == 0 {
		return nil
	} else if n == 1 {
		return lists[0]
	} else {
		l1 := mergeKLists(lists[0 : n/2])
		l2 := mergeKLists(lists[n/2:])
		return mergeList(l1, l2)
	}
}

func mergeList(l1, l2 *ListNode) *ListNode {
	dummy := &ListNode{Val: -1, Next: nil}
	p, p1, p2 := dummy, l1, l2
	for p1 != nil && p2 != nil {
		if p1.Val < p2.Val {
			p.Next = &ListNode{Val: p1.Val, Next: nil}
			p1 = p1.Next
		} else {
			p.Next = &ListNode{Val: p2.Val, Next: nil}
			p2 = p2.Next
		}
		p = p.Next
	}
	if p1 != nil {
		p.Next = p1
	} else if p2 != nil {
		p.Next = p2
	}
	return dummy.Next
}

// @lc code=end
