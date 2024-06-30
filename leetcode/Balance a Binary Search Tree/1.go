/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func balanceBST(root *TreeNode) *TreeNode {
	stack := []*TreeNode{}
	nodes := []*TreeNode{}

	stack = append(stack, root)
	for len(stack) > 0 {
		curr := stack[0]
		stack = stack[1:]
		if curr == nil {
			continue
		}
		nodes = append(nodes, curr)
		stack = append(stack, curr.Left)
		stack = append(stack, curr.Right)
	}

	sort.SliceStable(nodes, func(i, j int) bool {
		return nodes[i].Val < nodes[j].Val
	})

	return create(nodes)
}

func create(nodes []*TreeNode) *TreeNode {
	switch len(nodes) {
	case 0:
		return nil
	case 1:
		nodes[0].Left = nil
		nodes[0].Right = nil
		return nodes[0]
	}

	idx := len(nodes) / 2
	centerNode := nodes[idx]
	centerNode.Left = create(nodes[:idx])
	centerNode.Right = create(nodes[idx+1:])
	return centerNode
}
