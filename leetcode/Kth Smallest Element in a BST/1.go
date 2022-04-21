/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
var (
	count int
	res   *TreeNode
)

func inorder(root *TreeNode) {
	if root == nil {
		return
	}
	inorder(root.Left)
	count--
	if count == 0 {
		res = root
		return
	}
	inorder(root.Right)
}

func kthSmallest(root *TreeNode, k int) int {
	count = k
	res = nil
	inorder(root)
	if root == nil {
		return 0
	}
	return res.Val
}
