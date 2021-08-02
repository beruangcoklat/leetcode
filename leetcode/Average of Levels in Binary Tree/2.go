/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func averageOfLevels(root *TreeNode) []float64 {
	list := []*TreeNode{root}
	ret := []float64{}

	for len(list) > 0 {
		size := len(list)
		total := 0
		count := 0

		for i := 0; i < size; i++ {
			curr := list[0]

			list = list[1:]
			total += curr.Val
			count++
			if curr.Left != nil {
				list = append(list, curr.Left)
			}
			if curr.Right != nil {
				list = append(list, curr.Right)
			}
		}

		ret = append(ret, float64(total)/float64(count))
	}

	return ret
}
