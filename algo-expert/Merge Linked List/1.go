package main

// This is an input struct. Do not edit.
type LinkedList struct {
	Value int
	Next  *LinkedList
}

func MergeLinkedLists(headOne *LinkedList, headTwo *LinkedList) *LinkedList {
	var head, tail *LinkedList

	for headOne != nil || headTwo != nil {
		var curr *LinkedList
		if headTwo == nil {
			curr = headOne
			headOne = headOne.Next
		} else if headOne == nil {
			curr = headTwo
			headTwo = headTwo.Next
		} else if headOne.Value < headTwo.Value {
			curr = headOne
			headOne = headOne.Next
		} else {
			curr = headTwo
			headTwo = headTwo.Next
		}

		if head == nil {
			head = curr
			tail = curr
		} else {
			tail.Next = curr
			tail = curr
		}
	}
	tail.Next = nil

	return head
}
