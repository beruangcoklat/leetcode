/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
var res [][]int

func cloneList(list []*TreeNode) []*TreeNode {
	newList := make([]*TreeNode, len(list))
	for i, n := range list {
		newList[i] = n
	}
	return newList
}

func recur(root *TreeNode, list []*TreeNode, sum, targetSum int) {
	currSum := sum + root.Val

	if root.Left == nil && root.Right == nil {
		if currSum == targetSum {
			rows := []int{}
			for _, n := range list {
				rows = append(rows, n.Val)
			}
			rows = append(rows, root.Val)
			res = append(res, rows)

		}
		return
	}

	if root.Left != nil {
		recur(root.Left, append(cloneList(list), root), currSum, targetSum)
	}

	if root.Right != nil {
		recur(root.Right, append(cloneList(list), root), currSum, targetSum)
	}
}

func pathSum(root *TreeNode, targetSum int) [][]int {
	res = make([][]int, 0)
	if root == nil {
		return res
	}
	recur(root, []*TreeNode{}, 0, targetSum)
	return res
}
