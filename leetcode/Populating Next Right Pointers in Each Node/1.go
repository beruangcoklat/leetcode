/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */

func connect(root *Node) *Node {
	list := []*Node{root}

	for len(list) > 0 {
		size := len(list)

		for i := 0; i < len(list)-1; i++ {
			list[i].Next = list[i+1]
		}

		for i := 0; i < size; i++ {
			curr := list[0]
			list = list[1:]
			if curr == nil {
				continue
			}

			if curr.Left != nil {
				list = append(list, curr.Left)
			}
			if curr.Right != nil {
				list = append(list, curr.Right)
			}
		}
	}

	return root
}
