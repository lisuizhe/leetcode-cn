package leetcode

/*
 * @lc app=leetcode.cn id=19 lang=golang
 *
 * [19] 删除链表的倒数第 N 个结点
 */

// @lc code=start
//
// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Val: -1, Next: head}
	p := findNthFromEnd(dummy, n+1)
	p.Next = p.Next.Next
	return dummy.Next
}

func findNthFromEnd(head *ListNode, n int) *ListNode {
	p1, p2 := head, head
	for i := 0; i < n; i++ {
		p2 = p2.Next
	}
	for p2 != nil {
		p2 = p2.Next
		p1 = p1.Next
	}
	return p1
}

// @lc code=end
