package main

import "math"

type BST struct {
	Value int

	Left  *BST
	Right *BST
}

func (tree *BST) ValidateBst() bool {
	return Validate(tree, math.MinInt, math.MaxInt)
}

func Validate(tree *BST, min, max int) bool {
	if tree == nil {
		return true
	}

	if tree.Value < min || tree.Value >= max {
		return false
	}

	if !Validate(tree.Left, min, tree.Value) || !Validate(tree.Right, tree.Value, max) {
		return false
	}

	return true
}
