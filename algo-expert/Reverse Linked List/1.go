package main

type LinkedList struct {
	Value int
	Next  *LinkedList
}

func ReverseLinkedList(head *LinkedList) *LinkedList {
	curr := head
	var prev *LinkedList
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev
}
