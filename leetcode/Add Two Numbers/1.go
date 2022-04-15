/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func push(head, tail *ListNode, val int) (*ListNode, *ListNode) {
	if head == nil {
		head = &ListNode{Val: val}
		tail = head
	} else {
		tail.Next = &ListNode{Val: val}
		tail = tail.Next
	}
	return head, tail
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var head, tail *ListNode
	carrier := 0
	for l1 != nil || l2 != nil {
		a, b := 0, 0
		if l1 != nil {
			a = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			b = l2.Val
			l2 = l2.Next
		}

		total := a + b + carrier
		carrier = 0
		if total >= 10 {
			total -= 10
			carrier = 1
		}
		head, tail = push(head, tail, total)
	}

	if carrier > 0 {
		head, tail = push(head, tail, carrier)
	}
	return head
}
