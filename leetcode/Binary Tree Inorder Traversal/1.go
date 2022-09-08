/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	ret := []int{}
	l := inorderTraversal(root.Left)
	r := inorderTraversal(root.Right)

	ret = append(ret, l...)
	ret = append(ret, root.Val)
	ret = append(ret, r...)
	return ret
}
