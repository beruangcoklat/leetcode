/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func averageOfLevels(root *TreeNode) []float64 {
    ret := []float64{}
	list := []*TreeNode{root}

	for level := 0; len(list) > 0; level++ {
		size := len(list)
		total := float64(0)
		sizeNotNil := 0
		for i := 0; i < size; i++ {
			curr := list[0]
			list = list[1:]

			if curr == nil {
				continue
			}

			total += float64(curr.Val)
			sizeNotNil++

			list = append(list, curr.Left)
			list = append(list, curr.Right)
		}

		if sizeNotNil > 0 {
			ret = append(ret, total/float64(sizeNotNil))
		}
	}

	return ret
}