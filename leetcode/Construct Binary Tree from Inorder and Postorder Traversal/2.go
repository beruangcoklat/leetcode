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

func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(postorder) == 0 {
		return nil
	}

	root := &TreeNode{Val: postorder[len(postorder)-1]}

	if len(postorder) == 1 {
		return root
	}

	idx := 0
	for i, val := range inorder {
		if val == postorder[len(postorder)-1] {
			idx = i
			break
		}
	}

	leftInorder := inorder[:idx]
	rightInorder := inorder[idx+1:]
	root.Left = buildTree(leftInorder, getArr(postorder, leftInorder))
	root.Right = buildTree(rightInorder, getArr(postorder, rightInorder))

	return root
}

