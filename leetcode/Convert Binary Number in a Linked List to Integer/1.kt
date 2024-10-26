/**
 * Example:
 * var li = ListNode(5)
 * var v = li.`val`
 * Definition for singly-linked list.
 * class ListNode(var `val`: Int) {
 *     var next: ListNode? = null
 * }
 */
class Solution {
    fun getDecimalValue(head: ListNode?): Int {
        var curr = head
        val nums = mutableListOf<Int>()
        while (curr != null) {
            nums.add(curr.`val`)
            curr = curr.next
        }
        var ans = 0
        var exp = 0.0
        for (i in nums.size - 1 downTo 0) {
            if (nums[i] == 1) {
                ans += Math.pow(2.0, exp).toInt()
            }
            exp++
        }
        return ans
    }
}