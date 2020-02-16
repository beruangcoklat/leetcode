// https://leetcode.com/problems/add-two-numbers/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var curr, head *ListNode
	var send int = 0
	for l1 != nil && l2 != nil {
		total := l1.Val + l2.Val
		if send == 1 {
			total += 1
			send = 0
		}

		if total >= 10 {
			send = 1
			total -= 10
		}
		newnode := &ListNode{
			Val:  total,
			Next: nil,
		}

		if head == nil {
			head = newnode
			curr = head
		} else {
			curr.Next = newnode
			curr = newnode
		}

		l2 = l2.Next
		l1 = l1.Next
	}

	var l3 *ListNode
	if l1 != nil {
		l3 = l1
	}
	if l2 != nil {
		l3 = l2
	}

	for l3 != nil {
		if send == 1 {
			l3.Val += 1
			send = 0
		}
		if l3.Val >= 10 {
			l3.Val -= 10
			send = 1
		}
		curr.Next = l3
		curr = l3
		l3 = l3.Next
	}

	if send == 1 {
		curr.Next = &ListNode{
			Next: nil,
			Val:  1,
		}
	}
	return head
}