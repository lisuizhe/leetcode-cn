package leetcode

/*
 * @lc app=leetcode.cn id=142 lang=golang
 *
 * [142] 环形链表 II
 */

// @lc code=start
//
// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if slow == fast {
			break
		}
	}
	// no cycle
	if fast == nil || fast.Next == nil {
		return nil
	}
	// has cycle
	// assume:
	//  a: steps of head to start of loop
	//  b: steps of start of loop to meeting point
	//  c: steps of meeting point to start of loop
	//  n: # of rounds when fast meet slow in the loop
	// then:
	//  2(a+b) = a + n(b+c) + b
	// => a = n(b+c) - b
	// which means:
	//  from head to start of loop = from meeting point + n loop to start of loop
	slow = head
	for slow != fast {
		fast = fast.Next
		slow = slow.Next
	}
	return slow
}

// @lc code=end
