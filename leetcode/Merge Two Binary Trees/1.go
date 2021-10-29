/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func GetVal(t *TreeNode) int {
	if t == nil {
		return 0
	}
	return t.Val
}

func GetRight(t *TreeNode) *TreeNode {
	if t == nil {
		return nil
	}
	return t.Right
}

func GetLeft(t *TreeNode) *TreeNode {
	if t == nil {
		return nil
	}
	return t.Left
}

func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 == nil && root2 == nil {
		return nil
	}

	return &TreeNode{
		Val:   GetVal(root1) + GetVal(root2),
		Left:  mergeTrees(GetLeft(root1), GetLeft(root2)),
		Right: mergeTrees(GetRight(root1), GetRight(root2)),
	}
}
