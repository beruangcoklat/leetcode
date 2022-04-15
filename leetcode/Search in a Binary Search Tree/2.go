func searchBST(root *TreeNode, val int) *TreeNode {
	for root != nil {
		if root.Val == val {
			return root
		}
		if val > root.Val {
			root = root.Right
		} else {
			root = root.Left
		}
	}
	return nil
}
