import java.util.*;
import java.io.*;

class Node {
    Node left;
    Node right;
    int data;
    
    Node(int data) {
        this.data = data;
        left = null;
        right = null;
    }
}

class Solution {

	/* 
    
    class Node 
    	int data;
    	Node left;
    	Node right;
	*/
	public static void topView(Node root) {
        List<Node> list = new ArrayList<>();
        list.add(root);
        HashMap<Node, Integer> horizontalPos = new HashMap<>();
        horizontalPos.put(root, 0);

        int minPos = Integer.MAX_VALUE;
        int maxPos = Integer.MIN_VALUE;

        HashMap<Integer, Integer> mark = new HashMap<>();
        List<Integer> result = new ArrayList<>();

        while (!list.isEmpty()) {
            Node curr = list.get(0);
            list.remove(0);

            int currPos = horizontalPos.get(curr);
            if (!mark.containsKey(currPos)) {
                mark.put(currPos, curr.data);
                result.add(curr.data);

                minPos = Math.min(minPos, currPos);
                maxPos = Math.max(maxPos, currPos);
            }

            if (curr.right != null) {
                list.add(curr.right);
                horizontalPos.put(curr.right, currPos + 1);
            }
            if (curr.left != null) {
                list.add(curr.left);
                horizontalPos.put(curr.left, currPos - 1);
            }
        }

        for (int i = minPos; i <= maxPos; i++) {
            System.out.print(mark.get(i) + " ");
        }
    }

	public static Node insert(Node root, int data) {
        if(root == null) {
            return new Node(data);
        } else {
            Node cur;
            if(data <= root.data) {
                cur = insert(root.left, data);
                root.left = cur;
            } else {
                cur = insert(root.right, data);
                root.right = cur;
            }
            return root;
        }
    }

    public static void main(String[] args) {
        Scanner scan = new Scanner(System.in);
        int t = scan.nextInt();
        Node root = null;
        while(t-- > 0) {
            int data = scan.nextInt();
            root = insert(root, data);
        }
        scan.close();
        topView(root);
    }	
}
