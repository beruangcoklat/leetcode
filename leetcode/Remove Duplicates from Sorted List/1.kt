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
    fun deleteDuplicates(head: ListNode?): ListNode? {
        var curr = head
        while (curr != null) {
            var next = curr.next
            while (next != null) {
                if (curr.`val` == next.`val`) {
                    next = next.next
                } else {
                    curr.next = next
                    break
                }
            }
            if (next == null && curr.`val` == curr.next?.`val`) {
                curr.next = null
            }
            curr = curr.next
        }
        return head
    }
}