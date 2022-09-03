package main

func MinHeightBST(array []int) *BST {
	var tree *BST

	list := [][]int{array}
	for len(list) > 0 {
		curr := list[0]
		list = list[1:]

		if len(curr) <= 2 {
			for _, a := range curr {
				tree = tree.Insert(a)
			}
			continue
		}

		mid := len(curr) / 2
		tree = tree.Insert(curr[mid])
		list = append(list, curr[:mid], curr[mid+1:])
	}

	return tree
}

type BST struct {
	Value int

	Left  *BST
	Right *BST
}

func (tree *BST) Insert(value int) *BST {
	if tree == nil {
		return &BST{Value: value}
	}
	if value < tree.Value {
		if tree.Left == nil {
			tree.Left = &BST{Value: value}
		} else {
			tree.Left.Insert(value)
		}
	} else {
		if tree.Right == nil {
			tree.Right = &BST{Value: value}
		} else {
			tree.Right.Insert(value)
		}
	}
	return tree
}
