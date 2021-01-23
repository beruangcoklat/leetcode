// https://leetcode.com/problems/merge-two-sorted-lists/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func push(head, tail *ListNode, val int) (*ListNode, *ListNode) {
	if head == nil {
		c := &ListNode{Val: val, Next: nil}
		return c, c
	}

	tail.Next = &ListNode{Val: val, Next: nil}
	return head, tail.Next
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var head *ListNode = nil
	var tail *ListNode = nil

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			head, tail = push(head, tail, l1.Val)
			l1 = l1.Next
		} else {
			head, tail = push(head, tail, l2.Val)
			l2 = l2.Next
		}
	}

	for l1 != nil {
		head, tail = push(head, tail, l1.Val)
		l1 = l1.Next
	}

	for l2 != nil {
		head, tail = push(head, tail, l2.Val)
		l2 = l2.Next
	}

	return head
}