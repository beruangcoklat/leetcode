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
    
    ListNode getNextUntilNotLoop(ListNode node) {
        ListNode curr = node;
        while (curr != null) {
            if (curr.val != node.val) {
                return curr;
            }
            curr = curr.next;
        }
        return null;
    }

    public ListNode deleteDuplicates(ListNode head) {
        ListNode curr = head;
        ListNode prev = head;

        while (curr != null) {
            boolean isSame = curr.next != null && curr.val == curr.next.val;
            boolean isHead = curr == head;

            if (isSame && isHead) {
                head = getNextUntilNotLoop(head);
                curr = head;
                prev = head;
                continue;
            }

            if (isSame) {
                ListNode next = getNextUntilNotLoop(curr);
                prev.next = next;
                curr = next;
                continue;
            }

            prev = curr;
            curr = curr.next;
        }
        return head;
    }
    
}
