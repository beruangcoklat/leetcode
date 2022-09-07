/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func tree2str(root *TreeNode) string {
	if root == nil {
		return ""
	}

	left := ""
	right := ""

	if root.Left != nil {
		left = tree2str(root.Left)
	}
	if root.Right != nil {
		right = tree2str(root.Right)
	}

	if right != "" {
		left = fmt.Sprintf("(%s)", left)
		right = fmt.Sprintf("(%s)", right)
	} else if left != "" {
		left = fmt.Sprintf("(%s)", left)
	}

	return fmt.Sprintf("%d%s%s", root.Val, left, right)
}
