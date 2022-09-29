/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */

var visited map[int]*Node

func solve(node *Node) *Node {
	if node == nil {
		return nil
	}

	val, ok := visited[node.Val]
	if ok {
		return val
	}

	n := &Node{
		Val:       node.Val,
		Neighbors: nil,
	}
	visited[node.Val] = n

	for _, neighbor := range node.Neighbors {
		nn := solve(neighbor)
		if nn == nil {
			continue
		}
		n.Neighbors = append(n.Neighbors, nn)
	}

	return n
}

func cloneGraph(node *Node) *Node {
	visited = make(map[int]*Node)
	return solve(node)
}
