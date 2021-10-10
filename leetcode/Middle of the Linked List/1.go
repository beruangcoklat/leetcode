/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func middleNode(head *ListNode) *ListNode {
	curr := head
	count := 1

	for curr.Next != nil {
		curr = curr.Next
		count++
	}

	result := head
	for i := 0; i < (count / 2); i++ {
		result = result.Next
	}

	return result
}