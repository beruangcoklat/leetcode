/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
var nodes []*TreeNode

func bstToGst(root *TreeNode) *TreeNode {
	nodes = []*TreeNode{}
	inorder(root)
	for i := len(nodes) - 2; i >= 0; i-- {
		nodes[i].Val += nodes[i+1].Val
	}
	return root
}

func inorder(root *TreeNode) {
	if root == nil {
		return
	}
	inorder(root.Left)
	nodes = append(nodes, root)
	inorder(root.Right)
}
