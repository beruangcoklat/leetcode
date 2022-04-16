/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
var nodes []*TreeNode

func dfs(root *TreeNode) {
	if root == nil {
		return
	}
	dfs(root.Right)
	nodes = append(nodes, root)
	dfs(root.Left)
}

func convertBST(root *TreeNode) *TreeNode {
	nodes = make([]*TreeNode, 0)
	dfs(root)
	for i := 1; i < len(nodes); i++ {
		nodes[i].Val += nodes[i-1].Val
	}
	return root
}
