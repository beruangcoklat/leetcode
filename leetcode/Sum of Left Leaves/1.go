/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func solve(root *TreeNode, isLeftChild bool) int {
	if root == nil {
		return 0
	}
	if isLeftChild && root.Left == nil && root.Right == nil {
		return root.Val
	}
	return solve(root.Left, true) + solve(root.Right, false)
}

func sumOfLeftLeaves(root *TreeNode) int {
	return solve(root, false)
}
