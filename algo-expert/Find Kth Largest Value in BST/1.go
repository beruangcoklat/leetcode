package main

// This is an input class. Do not edit.
type BST struct {
	Value int

	Left  *BST
	Right *BST
}

var arr []int

func inorder(tree *BST) {
	if tree == nil {
		return
	}
	inorder(tree.Left)
	arr = append(arr, tree.Value)
	inorder(tree.Right)
}

func FindKthLargestValueInBst(tree *BST, k int) int {
	arr = []int{}
	inorder(tree)
	return arr[len(arr)-k]
}
