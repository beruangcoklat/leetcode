/**
 * Definition for a binary tree node.
 * class TreeNode(var `val`: Int = 0) {
 *     var left: TreeNode? = null
 *     var right: TreeNode? = null
 * }
 */
class Solution {
    fun lowestCommonAncestor(root: TreeNode?, p: TreeNode?, q: TreeNode?): TreeNode? {
        if (root == null) return null

        val pair = getMinMax(p!!.`val`, q!!.`val`)
        val min = pair.first
        val max = pair.second
        val rootVal = root?.`val` ?: 0

        if (rootVal in min..max) return root
        if (rootVal > max) return lowestCommonAncestor(root.left, p, q)
        return lowestCommonAncestor(root.right, p, q)
    }

    private fun getMinMax(a: Int, b: Int): Pair<Int, Int> =
        if (a < b) Pair(a, b) else Pair(b, a)

}