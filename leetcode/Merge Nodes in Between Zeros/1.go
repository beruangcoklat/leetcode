/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeNodes(head *ListNode) *ListNode {
	curr := head
	for curr != nil {
		if curr.Val == 0 {
			curr = curr.Next
			continue
		}
		sum := 0
		next := curr
		for next.Val != 0 {
			sum += next.Val
			next.Val = 0
			next = next.Next
		}
		curr.Val = sum
		curr = next
	}

	for head != nil && head.Val == 0 {
		head = head.Next
	}

	curr = head
	for curr != nil {
		next := curr.Next
		for next != nil && next.Val == 0 {
			next = next.Next
		}
		if curr == next {
			curr.Next = nil
			break
		}
		curr.Next = next
		curr = next
	}

	return head
}

