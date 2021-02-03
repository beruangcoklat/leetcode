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
    public List<List<Integer>> levelOrder(Node root) {
        List<List<Integer>> ret = new ArrayList<>();
        if(root == null) return ret;
        
        List<Node> list = new ArrayList<>();
        list.add(root);
        while (list.size() > 0) {
            int listSize = list.size();
            List<Integer> currLevel = new ArrayList<>();
            for (int i = 0; i < listSize; i++) {
                Node curr = list.get(0);
                list.remove(0);
                currLevel.add(curr.val);
                if (curr.children == null) continue;
                list.addAll(curr.children);
            }
            ret.add(currLevel);
        }
        return ret;
    }
}