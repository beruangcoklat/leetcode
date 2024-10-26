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
    fun sortedArrayToBST(nums: IntArray): TreeNode? {
        when (nums.size) {
            0 -> return null
            1 -> return TreeNode(nums[0])
        }
        val mid = nums.size / 2
        val root = TreeNode(nums[mid])
        root.left = sortedArrayToBST(nums.slice(0..mid - 1).toIntArray())
        root.right = sortedArrayToBST(nums.slice(mid + 1..nums.size - 1).toIntArray())
        return root
    }
}