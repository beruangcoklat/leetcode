/*
// Definition for a Node.
class Node {
    public int val;
    public List<Node> children;

    public Node() {}

    public Node(int _val) {
        val = _val;
    }

    public Node(int _val, List<Node> _children) {
        val = _val;
        children = _children;
    }
};
*/

class Solution {
    List<Integer> list;
    
    void postorderRec(Node root){
        if(root == null) return;
        
        for(Node c : root.children){
            postorderRec(c);
        }
        
        list.add(root.val);
    }
    
    public List<Integer> postorder(Node root) {
        list = new ArrayList();
        postorderRec(root);
        return list;
    }
}