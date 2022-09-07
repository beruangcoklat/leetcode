/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
var parentMap map[*TreeNode]*TreeNode
var result int

func preorder(node, parent *TreeNode) {
	if node == nil {
		return
	}

	parentMap[node] = parent

	curr := node
	x := curr.Val
	isGood := true
	for curr != nil {
		if curr.Val > x {
			isGood = false
			break
		}
		curr = parentMap[curr]
	}

	if isGood {
		result++
	}

	preorder(node.Left, node)
	preorder(node.Right, node)
}

func goodNodes(root *TreeNode) int {
	parentMap = make(map[*TreeNode]*TreeNode)
	result = 0
	preorder(root, nil)
	return result
}
