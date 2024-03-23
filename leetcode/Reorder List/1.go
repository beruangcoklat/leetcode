/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reorderList(head *ListNode) {
	arr := []*ListNode{}

	curr := head
	for curr != nil {
		arr = append(arr, curr)
		curr = curr.Next
	}

	isLeft := false
	size := len(arr)

	leftIdx := 0
	rightIdx := size - 1
	for i := 0; i < size; i++ {
		isLeft = !isLeft

		var from, target int
		if isLeft {
			from = leftIdx
			target = rightIdx
			leftIdx++
		} else {
			from = rightIdx
			target = leftIdx
			rightIdx--
		}

		if i == size-1 {
			arr[from].Next = nil
		} else {
			arr[from].Next = arr[target]
		}
	}
}