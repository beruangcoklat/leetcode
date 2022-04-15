/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findReplacer(root *TreeNode, low, high int) *TreeNode {
	for root != nil && (root.Val > high || root.Val < low) {
		if root.Val > high {
			root = root.Left
		} else {
			root = root.Right
		}
	}
	return root
}

func trimBST(root *TreeNode, low int, high int) *TreeNode {
	root = findReplacer(root, low, high)
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]
		if curr == nil {
			continue
		}
		curr.Right = findReplacer(curr.Right, low, high)
		curr.Left = findReplacer(curr.Left, low, high)
		queue = append(queue, curr.Right, curr.Left)
	}

	return root
}