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
    public ListNode mergeTwoLists(ListNode l1, ListNode l2) {
        ListNode curr = null, head = null;
        while(l1 != null && l2 != null){
            int min = 0;
            if(l1.val <= l2.val){
                min = l1.val;
                l1 = l1.next;
            }else{
                min = l2.val;
                l2 = l2.next;
            }
            
            ListNode newNode = new ListNode(min);
            if(curr == null){
                curr = newNode;
                head = curr;
            }else{
                curr.next = newNode;
                curr = curr.next;
            }
        }
        
        while(l1 != null){
            int min = l1.val;
            l1 = l1.next;
            
            ListNode newNode = new ListNode(min);
            if(curr == null){
                curr = newNode;
                head = curr;
            }else{
                curr.next = newNode;
                curr = curr.next;
            }
        }
        while(l2 != null){
            int min = l2.val;
            l2 = l2.next;
            
            ListNode newNode = new ListNode(min);
            if(curr == null){
                curr = newNode;
                head = curr;
            }else{
                curr.next = newNode;
                curr = curr.next;
            }
        }
        
        return head;
    }
}
