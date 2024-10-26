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
    private var ans: MutableList<Int> = mutableListOf()

    fun getAllElements(root1: TreeNode?, root2: TreeNode?): List<Int> {
        ans = mutableListOf()
        preorder(root1)
        preorder(root2)
        ans.sort()
        return ans
    }

    private fun preorder(root: TreeNode?) {
        if (root == null) return
        ans.add(root.`val`)
        preorder(root.left)
        preorder(root.right)
    }
}