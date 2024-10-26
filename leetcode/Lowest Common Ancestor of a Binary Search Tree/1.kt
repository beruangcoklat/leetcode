/**
 * Definition for a binary tree node.
 * class TreeNode(var `val`: Int = 0) {
 *     var left: TreeNode? = null
 *     var right: TreeNode? = null
 * }
 */
class Solution {
    private var parentMap: MutableMap<Int, TreeNode> = mutableMapOf()

    fun lowestCommonAncestor(root: TreeNode?, p: TreeNode?, q: TreeNode?): TreeNode? {
        parentMap = mutableMapOf()
        setParentMap(root)


        val m: MutableMap<Int, Boolean> = mutableMapOf()

        var curr = p
        while (curr != null) {
            m[curr.`val`] = true
            curr = parentMap[curr.`val`]
        }

        curr = q
        while (curr != null) {
            if (m[curr.`val`] != null) {
                return curr
            }
            curr = parentMap[curr.`val`]
        }

        return null
    }

    private fun setParentMap(root: TreeNode?) {
        if (root == null) return
        parentMap[root.left?.`val` ?: -1] = root
        parentMap[root.right?.`val` ?: -1] = root
        setParentMap(root.left)
        setParentMap(root.right)
    }
}