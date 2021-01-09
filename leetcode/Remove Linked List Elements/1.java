/**
 * Definition for singly-linked list.
 * public class ListNode {
 *     int val;
 *     ListNode next;
 *     ListNode() {}
 *     ListNode(int val) { this.val = val; }
 *     ListNode(int val, ListNode next) { this.val = val; this.next = next; }
 * }
 */
class Solution {
    public ListNode removeElements(ListNode head, int val) {
        ListNode curr = head;
        ListNode prev = head;

        while (curr != null) {
            if (curr.val != val) {
                prev = curr;
                curr = curr.next;
                continue;
            }
            boolean isHead = curr == head;
            boolean isTail = curr.next == null;

            if (isHead && isTail) {
                return null;
            }
            
            if (!isHead && !isTail) {
                prev.next = curr.next;
                curr = prev.next;
                continue;
            }

            if (isHead) {
                head = head.next;
                curr = head;
                prev = curr;
                continue;
            }

            if (isTail) {
                prev.next = null;
                return head;
            }
        }
        return head;
    }
}
