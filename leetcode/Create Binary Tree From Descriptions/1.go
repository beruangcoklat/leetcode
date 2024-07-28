/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func createBinaryTree(descriptions [][]int) *TreeNode {
	nodeMap := make(map[int]*TreeNode)
	for _, description := range descriptions {
		parent := description[0]
		child := description[1]
		isLeft := description[2] == 1

		if _, ok := nodeMap[parent]; !ok {
			nodeMap[parent] = &TreeNode{Val: parent}
		}
		if _, ok := nodeMap[child]; !ok {
			nodeMap[child] = &TreeNode{Val: child}
		}
		if isLeft {
			nodeMap[parent].Left = nodeMap[child]
		} else {
			nodeMap[parent].Right = nodeMap[child]
		}
	}

	// find node without parent
	childMap := make(map[int]struct{})
	for _, description := range descriptions {
		child := description[1]
		childMap[child] = struct{}{}
	}

	for _, node := range nodeMap {
		if _, ok := childMap[node.Val]; !ok {
			return node
		}
	}

	return nil
}