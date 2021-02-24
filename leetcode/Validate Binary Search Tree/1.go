/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func solve(root *TreeNode, min, max int, checkMin, checkMax bool) bool {
	if root == nil {
		return true
	}

	if checkMin && root.Val <= min {
		return false
	}

	if checkMax && root.Val >= max {
		return false
	}

	return solve(root.Right, root.Val, max, true, checkMax) && solve(root.Left, min, root.Val, checkMin, true)
}

func isValidBST(root *TreeNode) bool {
	return solve(root, 0, 0, false, false)
}
