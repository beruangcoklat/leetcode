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
    fun bstFromPreorder(preorder: IntArray): TreeNode? {
        var root: TreeNode? = null
        preorder.forEach { n ->
            root = insert(root, n)
        }
        return root
    }

    private fun insert(root: TreeNode?, value: Int): TreeNode? {
        if (root == null) return TreeNode(value)
        if (value < root.`val`) root.left = insert(root.left, value)
        else root.right = insert(root.right, value)
        return root
    }
}