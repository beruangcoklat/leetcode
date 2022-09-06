/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
var nodeMap map[int][]*TreeNode
var maxDepth int

func preorder(root *TreeNode, depth int) {
	if root == nil {
		return
	}

	_, ok := nodeMap[depth]
	if !ok {
		nodeMap[depth] = []*TreeNode{root}
	} else {
		nodeMap[depth] = append(nodeMap[depth], root)
	}

	if depth > maxDepth {
		maxDepth = depth
	}

	preorder(root.Left, depth+1)
	preorder(root.Right, depth+1)
}

func prune(root *TreeNode, deletedNodeMap map[*TreeNode]struct{}) *TreeNode {
	if root == nil {
		return nil
	}
	_, deleted := deletedNodeMap[root]
	if deleted {
		return nil
	}
	root.Left = prune(root.Left, deletedNodeMap)
	root.Right = prune(root.Right, deletedNodeMap)
	return root
}

func pruneTree(root *TreeNode) *TreeNode {
	nodeMap = make(map[int][]*TreeNode)
	preorder(root, 0)
	deletedNodeMap := make(map[*TreeNode]struct{})

	for depth := maxDepth; depth >= 0; depth-- {
		nodes := nodeMap[depth]
		for _, node := range nodes {
			left := false
			right := false

			if node.Left != nil {
				_, left = deletedNodeMap[node.Left]
			} else {
				left = true
			}

			if node.Right != nil {
				_, right = deletedNodeMap[node.Right]
			} else {
				right = true
			}

			if node.Val != 1 && left && right {
				deletedNodeMap[node] = struct{}{}
			}
		}
	}

	return prune(root, deletedNodeMap)
}
