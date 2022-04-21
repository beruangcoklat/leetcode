/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
var arr []*TreeNode

func inorder(root *TreeNode) {
	if root == nil {
		return
	}
	inorder(root.Left)
	arr = append(arr, root)
	inorder(root.Right)
}

func recoverTree(root *TreeNode) {
	arr = []*TreeNode{}
	inorder(root)
	a := -1
	b := -1

	for i := 0; i < len(arr); i++ {
		isWeird := (i+1 < len(arr) && arr[i].Val > arr[i+1].Val) ||
			(i-1 >= 0 && arr[i-1].Val > arr[i].Val)

		if !isWeird {
			continue
		}

		if a == -1 {
			a = i
			continue
		}

		b = i
	}
	arr[b].Val, arr[a].Val = arr[a].Val, arr[b].Val
}
