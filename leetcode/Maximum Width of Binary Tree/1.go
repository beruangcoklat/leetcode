/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

var table map[int][]int

func dfs(root *TreeNode, depth, pos int) {
	if root == nil {
		return
	}

	if _, ok := table[depth]; !ok {
		table[depth] = []int{}
	}

	table[depth] = append(table[depth], pos)

	dfs(root.Left, depth+1, pos*2)
	dfs(root.Right, depth+1, pos*2+1)
}

func widthOfBinaryTree(root *TreeNode) int {
	table = make(map[int][]int)
	dfs(root, 0, 1)

	ret := 1
	r := 1
	for depth, row := range table {
		if depth == 0 {
			continue
		}

		r *= 2
		max, min := row[0], row[0]

		for i := 1; i < len(row); i++ {
			if row[i] > max {
				max = row[i]
			}
			if row[i] < min {
				min = row[i]
			}
		}
		curr := max - min + 1
		if curr > ret {
			ret = curr
		}
	}
	return ret
}
