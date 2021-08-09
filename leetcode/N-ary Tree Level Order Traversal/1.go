/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func levelOrder(root *Node) [][]int {
	result := [][]int{}

	if root == nil {
		return result
	}

	list := []*Node{root}
	for len(list) > 0 {
		size := len(list)
		rows := []int{}
		for i := 0; i < size; i++ {
			curr := list[0]
			list = list[1:]
			rows = append(rows, curr.Val)
			for _, child := range curr.Children {
				list = append(list, child)
			}
		}
		result = append(result, rows)
	}
	return result
}
