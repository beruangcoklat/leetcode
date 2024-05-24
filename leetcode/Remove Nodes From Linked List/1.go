/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNodes(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	arr := []int{}
	curr := head
	for curr != nil {
		arr = append(arr, curr.Val)
		curr = curr.Next
	}

	maxFromRight := make([]int, len(arr))
	maxFromRight[len(arr)-1] = arr[len(arr)-1]
	for i := len(arr) - 2; i >= 0; i-- {
		if arr[i] > maxFromRight[i+1] {
			maxFromRight[i] = arr[i]
		} else {
			maxFromRight[i] = maxFromRight[i+1]
		}
	}

	curr = head
	idx := 0
	validNode := []*ListNode{}
	for curr != nil {
		if curr.Val >= maxFromRight[idx] {
			validNode = append(validNode, curr)
		}
		curr = curr.Next
		idx++
	}

	for i := 0; i < len(validNode)-1; i++ {
		validNode[i].Next = validNode[i+1]
	}
	validNode[len(validNode)-1].Next = nil
	return validNode[0]
}
