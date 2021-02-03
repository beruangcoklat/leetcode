/**
 * Definition for a binary tree node.
 * public class TreeNode {
 *     int val;
 *     TreeNode left;
 *     TreeNode right;
 *     TreeNode() {}
 *     TreeNode(int val) { this.val = val; }
 *     TreeNode(int val, TreeNode left, TreeNode right) {
 *         this.val = val;
 *         this.left = left;
 *         this.right = right;
 *     }
 * }
 */
class Solution {
    
    List<Integer> list;
    HashMap<Integer, Boolean> map;
    
    void inorder(TreeNode root){
        if(root == null) return;
        inorder(root.left);
        
        if(!map.containsKey(root.val)){
            map.put(root.val, true);
            list.add(root.val);
        }

        inorder(root.right);
    }

    public int findSecondMinimumValue(TreeNode root) {
        list = new ArrayList<>();
        map = new HashMap<>();
        inorder(root);
        Collections.sort(list);
        return list.size() < 2 ? -1 : list.get(1);
    }
}