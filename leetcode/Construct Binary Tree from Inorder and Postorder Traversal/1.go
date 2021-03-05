/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(inorder []int, postorder []int) *TreeNode {
    if len(postorder) == 0 {
		return nil
	}

	if len(postorder) == 1 {
		return &TreeNode{postorder[0], nil, nil}
	}

    root := &TreeNode{postorder[len(postorder) - 1], nil, nil}
	var idxRoot int
	for i, v := range inorder {
		if v == root.Val {
			idxRoot = i
			break
		}
	}

	root.Left = buildTree(inorder[:idxRoot], postorder[:idxRoot])
    root.Right = buildTree(inorder[idxRoot+1:], postorder[idxRoot:len(postorder) - 1])

	return root
}