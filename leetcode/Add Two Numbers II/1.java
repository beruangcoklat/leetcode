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
    
    String add(String num1, String num2) {
        String ret = "";
        int idx1 = num1.length() - 1;
        int idx2 = num2.length() - 1;
        int carrier = 0;
        while (idx1 >= 0 && idx2 >= 0) {
            int a = num1.charAt(idx1) - '0';
            int b = num2.charAt(idx2) - '0';

            int total = a + b + carrier;
            carrier = 0;
            if (total > 9) {
                carrier = 1;
                total -= 10;
            }

            ret = (total + "") + ret;

            idx1--;
            idx2--;
        }

        while (idx1 >= 0) {
            int a = num1.charAt(idx1) - '0';
            int total = a + carrier;
            carrier = 0;
            if (total > 9) {
                carrier = 1;
                total -= 10;
            }
            ret = (total + "") + ret;
            idx1--;
        }

        while (idx2 >= 0) {
            int a = num2.charAt(idx2) - '0';
            int total = a + carrier;
            carrier = 0;
            if (total > 9) {
                carrier = 1;
                total -= 10;
            }
            ret = (total + "") + ret;
            idx2--;
        }

        return (carrier == 1 ? "1" : "") + ret;
    }
    
    public ListNode addTwoNumbers(ListNode l1, ListNode l2) {
        String num1 = "";
        String num2 = "";
        while(l1 != null && l2 != null){
            num1 += (l1.val + "");
            num2 += (l2.val + "");
            l1 = l1.next;
            l2 = l2.next;
        }
        
        while(l1 != null){
            num1 += (l1.val + "");
            l1 = l1.next;
        }
        
        while(l2 != null){
            num2 += (l2.val + "");
            l2 = l2.next;
        }
        
        String str = add(num1, num2);
        
        ListNode head = null;
        ListNode tail = null;
        for(int i=0 ; i < str.length() ; i++){
            ListNode newNode = new ListNode(str.charAt(i) - '0');
            if(head == null){
                head = newNode;
                tail = newNode;
                continue;
            }
            
            tail.next = newNode;
            tail = tail.next;
        }
        
        return head;
    }
}
