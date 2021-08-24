/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
var arr []int

func recur(root *TreeNode) {
	if root == nil {
		return
	}
	arr = append(arr, root.Val)
	recur(root.Left)
	recur(root.Right)
}

func findTarget(root *TreeNode, k int) bool {
	arr = []int{}
	recur(root)

	dict := map[int]struct{}{}
	for i := 0; i < len(arr); i++ {
		_, ok := dict[k-arr[i]]
		if ok {
			return true
		}
		dict[arr[i]] = struct{}{}
	}

	return false
}
