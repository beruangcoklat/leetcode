package main

type LinkedList struct {
	Value int
	Next  *LinkedList
}

func ReverseLinkedList(head *LinkedList) *LinkedList {
	prevMap := make(map[*LinkedList]*LinkedList)

	tail := head
	for tail.Next != nil {
		tail = tail.Next
	}

	var prev *LinkedList
	curr := head
	for curr != nil {
		prevMap[curr] = prev
		prev = curr
		curr = curr.Next
	}

	curr = tail
	for curr != nil {
		curr.Next = prevMap[curr]
		curr = curr.Next
	}

	return tail
}
