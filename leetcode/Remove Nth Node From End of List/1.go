/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	curr := head
	count := 1
	for curr.Next != nil {
		count++
		curr = curr.Next
	}

	if count == 1 {
		return nil
	}

	removedNode := count - n + 1
	if removedNode == 1 {
		return head.Next
	}

	i := 1
	curr = head
	prev := head
	for curr != nil {
		if i != removedNode {
			prev = curr
			curr = curr.Next
			i++
			continue
		}

		prev.Next = curr.Next
		curr = nil
		return head
	}

	return nil
}