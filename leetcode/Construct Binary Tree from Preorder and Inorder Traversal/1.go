/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	if len(preorder) == 1 {
		return &TreeNode{preorder[0], nil, nil}
	}

	root := &TreeNode{preorder[0], nil, nil}
	var idxRoot int
	for i, v := range inorder {
		if v == root.Val {
			idxRoot = i
			break
		}
	}

	root.Left = buildTree(preorder[1:idxRoot+1], inorder[:idxRoot])
	root.Right = buildTree(preorder[idxRoot+1:], inorder[idxRoot+1:])

	return root
}