/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSame(root1 *TreeNode, root2 *TreeNode) bool {
	if root1 == nil && root2 == nil {
		return true
	}
	if root1 == nil || root2 == nil || root1.Val != root2.Val {
		return false
	}
	return isSame(root1.Left, root2.Left) && isSame(root1.Right, root2.Right)
}

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	if isSame(root, subRoot) {
		return true
	}
	if root.Left != nil && isSubtree(root.Left, subRoot) {
		return true
	}
	if root.Right != nil && isSubtree(root.Right, subRoot) {
		return true
	}
	return false
}
