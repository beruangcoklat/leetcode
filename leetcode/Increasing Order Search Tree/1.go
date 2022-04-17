/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
var nodes []*TreeNode

func inorder(root *TreeNode) {
	if root == nil {
		return
	}
	inorder(root.Left)
	nodes = append(nodes, root)
	inorder(root.Right)
}

func increasingBST(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	nodes = make([]*TreeNode, 0)
	inorder(root)
	for i := 0; i < len(nodes); i++ {
		nodes[i].Left = nil
		if i+1 < len(nodes) {
			nodes[i].Right = nodes[i+1]
		} else {
			nodes[i].Right = nil
		}
	}
	return nodes[0]
}
