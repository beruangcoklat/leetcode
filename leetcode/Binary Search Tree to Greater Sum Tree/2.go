/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func bstToGst(root *TreeNode) *TreeNode {
	stack := []*TreeNode{root}
	nodes := []*TreeNode{}
	for len(stack) > 0 {
		curr := stack[0]
		stack = stack[1:]
		if curr == nil {
			continue
		}
		nodes = append(nodes, curr)
		stack = append(stack, curr.Left, curr.Right)
	}

	sort.SliceStable(nodes, func(i, j int) bool {
		return nodes[i].Val < nodes[j].Val
	})

	for i := len(nodes) - 2; i >= 0; i-- {
		nodes[i].Val += nodes[i+1].Val
	}
	return root
}