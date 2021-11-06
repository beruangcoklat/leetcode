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

func constructFromPrePost(preorder []int, postorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	root := &TreeNode{Val: preorder[0]}

	if len(preorder) == 1 {
		return root
	}

	preorder = preorder[1:]
	postorder = postorder[:len(postorder)-1]

	idx := 0
	for i, val := range postorder {
		if val == preorder[0] {
			idx = i
			break
		}
	}

	leftPostorder := postorder[:idx+1]
	rightPostorder := postorder[idx+1:]
	root.Left = constructFromPrePost(getArr(preorder, leftPostorder), leftPostorder)
	root.Right = constructFromPrePost(getArr(preorder, rightPostorder), rightPostorder)

	return root
}
