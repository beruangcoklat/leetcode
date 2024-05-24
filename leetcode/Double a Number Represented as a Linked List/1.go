/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func doubleIt(head *ListNode) *ListNode {
	curr := head
	var prev, tail *ListNode
	prevMap := make(map[*ListNode]*ListNode)

	for curr != nil {
		tail = curr
		prevMap[curr] = prev
		prev = curr
		curr = curr.Next
	}

	carrier := 0
	for tail != nil {
		tail.Val = (tail.Val * 2) + carrier
		carrier = 0

		if tail.Val >= 10 {
			tail.Val -= 10
			carrier = 1
		}

		tail = prevMap[tail]
	}

	if carrier == 1 {
		return &ListNode{
			Val:  1,
			Next: head,
		}
	}
	return head
}
