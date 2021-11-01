/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
    if head == nil {
        return nil
    }
	arr := []*ListNode{}
	curr := head
	for curr != nil {
		arr = append(arr, curr)
		curr = curr.Next
	}
	for i := len(arr) - 1; i >= 1; i-- {
		arr[i].Next = arr[i-1]
	}
	arr[0].Next = nil
	return arr[len(arr)-1]
}
