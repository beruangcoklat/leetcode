/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeInBetween(list1 *ListNode, a int, b int, list2 *ListNode) *ListNode {
	curr := list1
	prev := curr
	for i := 0; i < a; i++ {
		prev = curr
		curr = curr.Next
	}

	prevStart := prev

	for i := 0; i < b-a; i++ {
		prev = curr
		curr = curr.Next
	}

	end := curr
	prevStart.Next = list2

	for list2.Next != nil {
		list2 = list2.Next
	}
	list2.Next = end.Next
	return list1
}