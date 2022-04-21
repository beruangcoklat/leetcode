/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type BSTIterator struct {
	arr []*TreeNode
	idx int
}

func (this *BSTIterator) inorder(root *TreeNode) {
	if root == nil {
		return
	}
	this.inorder(root.Left)
	this.arr = append(this.arr, root)
	this.inorder(root.Right)
}

func Constructor(root *TreeNode) BSTIterator {
	b := BSTIterator{
		arr: []*TreeNode{},
	}
	b.inorder(root)
	return b
}

func (this *BSTIterator) Next() int {
	ret := this.arr[this.idx].Val
	this.idx++
	return ret
}

func (this *BSTIterator) HasNext() bool {
	return this.idx < len(this.arr)
}

/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */
