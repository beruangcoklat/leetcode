/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {
	var head, res *ListNode

	for true {
		minIdx := -1
		for i, n := range lists {
			if n == nil {
				continue
			}

			if minIdx == -1 || n.Val < lists[minIdx].Val {
				minIdx = i
			}
		}

		if minIdx == -1 {
			break
		}

		if res == nil {
			res = lists[minIdx]
			head = res
		} else {
			res.Next = lists[minIdx]
			res = res.Next
		}

		lists[minIdx] = lists[minIdx].Next
	}

	return head
}
