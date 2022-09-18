package leetcode

/*
 * @lc app=leetcode.cn id=86 lang=golang
 *
 * [86] 分隔链表
 */

// @lc code=start
//
// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
	if head == nil {
		return nil
	}
	dummyLeft := &ListNode{Val: -1, Next: nil}
	dummyRight := &ListNode{Val: -1, Next: nil}
	pLeft := dummyLeft
	pRight := dummyRight
	p := head
	for p != nil {
		if p.Val < x {
			pLeft.Next = &ListNode{Val: p.Val, Next: nil}
			pLeft = pLeft.Next
		} else {
			pRight.Next = &ListNode{Val: p.Val, Next: nil}
			pRight = pRight.Next
		}
		p = p.Next
	}
	if dummyLeft.Next == nil {
		return dummyRight.Next
	} else {
		pLeft.Next = dummyRight.Next
	}
	return dummyLeft.Next
}

// @lc code=end
