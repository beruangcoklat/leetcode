/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
	counter := 0
	list := []*ListNode{}
	ret := head
	first := true

	var prev *ListNode
	curr := head
	prevMap := make(map[*ListNode]*ListNode)
	for curr != nil {
		prevMap[curr] = prev
		prev = curr
		curr = curr.Next
	}

	for head != nil {
		list = append(list, head)
		head = head.Next
		counter++

		if counter < k {
			continue
		}

		lastIdx := len(list) - 1
		if first {
			ret = list[lastIdx]
			first = false
		}

		mostRightFromPrevGroup := prevMap[list[0]]
		mostLeftNextGroup := list[lastIdx].Next

		for i := len(list) - 1; i > 0; i-- {
			list[i].Next = list[i-1]
			prevMap[list[i-1]] = list[i]
		}
		list[0].Next = mostLeftNextGroup
		prevMap[list[lastIdx]] = prev

		if mostRightFromPrevGroup != nil {
			mostRightFromPrevGroup.Next = list[lastIdx]
		}
		prevMap[mostLeftNextGroup] = list[0]
		list = []*ListNode{}
		counter = 0
	}

	return ret
}
