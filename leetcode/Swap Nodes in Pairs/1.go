/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
	curr := head
	head = nil
	var left *ListNode

	for curr != nil {
		right := curr.Next
		if right == nil {
			if head == nil {
				return curr
			}
			return head
		}

		if left != nil {
			left.Next = right
		}

		curr.Next = right.Next
		right.Next = curr

		if head == nil {
			head = right
		}

		left = curr
		curr = curr.Next
	}

	return head
}
