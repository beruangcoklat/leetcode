/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type data struct {
	val int
	row int
	col int
}

var datas []data

func traverse(root *TreeNode, row, col int) {
	if root == nil {
		return
	}

	datas = append(datas, data{
		val: root.Val,
		row: row,
		col: col,
	})
	traverse(root.Left, row+1, col-1)
	traverse(root.Right, row+1, col+1)
}

func verticalTraversal(root *TreeNode) [][]int {
	datas = []data{}
	traverse(root, 0, 0)

	sort.SliceStable(datas, func(i, j int) bool {
		a, b := datas[i], datas[j]

		if a.col < b.col {
			return true
		}

		if a.col == b.col {
			if a.row < b.row {
				return true
			}
			if a.row == b.row {
				return a.val < b.val
			}
			return false
		}

		return false
	})

	ret := [][]int{}
	idx := -1
	prevCol := math.MaxInt
	for _, data := range datas {
		if data.col == prevCol {
			ret[idx] = append(ret[idx], data.val)
		} else {
			idx++
			ret = append(ret, []int{data.val})
		}
		prevCol = data.col
	}

	return ret
}
