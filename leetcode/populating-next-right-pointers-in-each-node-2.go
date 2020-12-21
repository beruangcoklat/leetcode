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
    if root == nil{
        return nil
    }
    curr := root
    list := []*Node{curr}
    
    for len(list) > 0 {
        nodes := []*Node{}

        for _, n := range list {
            if n.Left != nil{
                nodes = append(nodes, n.Left)
            }
            if n.Right != nil{
                nodes = append(nodes, n.Right)
            }
        }
        
        list = nodes

        for i := 0 ; i<len(nodes) ; i++{
            if i == len(nodes)-1{
                nodes[i].Next = nil
                continue
            }
            nodes[i].Next = nodes[i+1]
        }        
    }
    return root
}
