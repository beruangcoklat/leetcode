/*
// Definition for a Node.
class Node {
    int val;
    Node next;
    Node random;

    public Node(int val) {
        this.val = val;
        this.next = null;
        this.random = null;
    }
}
*/

class Solution {
    class Pair {
        Node head, tail;
        Pair(Node head, Node tail){
            this.head=head;
            this.tail=tail;
        }
    }

    Pair push(Node head, Node tail, int val) {
        Node newNode = new Node(val);
        if (head == null) {
            head = newNode;
            tail = newNode;
        } else {
            tail.next = newNode;
            tail = tail.next;
        }
        return new Pair(head, tail);
    }
    
    public Node copyRandomList(Node oriHead) {
        Node head = null;
        Node tail = null;
        Node currOri = oriHead;

        HashMap<Node, Integer> mapOldNodeToPos = new HashMap<>();
        HashMap<Integer, Node> mapPosToNewNode = new HashMap<>();

        int i = 0;
        while (currOri != null) {
            mapOldNodeToPos.put(currOri, i);
            Pair r = push(head, tail, currOri.val);
            head = r.head;
            tail = r.tail;
            currOri = currOri.next;
            mapPosToNewNode.put(i, tail);
            i++;
        }

        currOri = oriHead;
        Node curr = head;
        while (curr != null) {
            if (currOri.random != null) {
                int pos = mapOldNodeToPos.get(currOri.random);
                curr.random = mapPosToNewNode.get(pos);
            }

            currOri = currOri.next;
            curr = curr.next;
        }

        return head;
    }
}
