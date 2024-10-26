/**
 * Example:
 * var ti = TreeNode(5)
 * var v = ti.`val`
 * Definition for a binary tree node.
 * class TreeNode(var `val`: Int) {
 *     var left: TreeNode? = null
 *     var right: TreeNode? = null
 * }
 */
class Solution {
    private var list1: MutableList<Int> = mutableListOf()
    private var list2: MutableList<Int> = mutableListOf()

    fun getAllElements(root1: TreeNode?, root2: TreeNode?): List<Int> {
        list1 = mutableListOf()
        list2 = mutableListOf()

        inorder(root1, list1)
        inorder(root2, list2)

        val ans: MutableList<Int> = mutableListOf()

        var idx1 = 0
        var idx2 = 0
        while (idx1 < list1.size && idx2 < list2.size) {
            if (list1[idx1] < list2[idx2]) {
                ans.add(list1[idx1])
                idx1++
            } else {
                ans.add(list2[idx2])
                idx2++
            }
        }
        while (idx1 < list1.size) {
            ans.add(list1[idx1])
            idx1++
        }
        while (idx2 < list2.size) {
            ans.add(list2[idx2])
            idx2++
        }
        return ans
    }

    private fun inorder(root: TreeNode?, list: MutableList<Int>) {
        if (root == null) return
        inorder(root.left, list)
        list.add(root.`val`)
        inorder(root.right, list)
    }
}