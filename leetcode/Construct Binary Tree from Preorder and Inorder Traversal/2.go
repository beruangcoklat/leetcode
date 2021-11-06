/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func getArr(arr []int, available []int) []int {
	availableMap := map[int]struct{}{}
	for _, a := range available {
		availableMap[a] = struct{}{}
	}

	newArr := []int{}
	for _, a := range arr {
		if _, ok := availableMap[a]; !ok {
			continue
		}
		newArr = append(newArr, a)
	}

	return newArr
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	root := &TreeNode{Val: preorder[0]}

	if len(preorder) == 1 {
		return root
	}

	idx := 0
	for i, val := range inorder {
		if val == preorder[0] {
			idx = i
			break
		}
	}

	leftInorder := inorder[:idx]
	rightInorder := inorder[idx+1:]
	root.Left = buildTree(getArr(preorder, leftInorder), leftInorder)
	root.Right = buildTree(getArr(preorder, rightInorder), rightInorder)

	return root
}
