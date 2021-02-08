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
    
    boolean isSame(TreeNode a, TreeNode b){
        if(a == null && b == null) return true;
        if(a == null && b != null) return false;
        if(a != null && b == null) return false;
        return a.val == b.val;
    }
    
    boolean check(TreeNode a, TreeNode b){
        if(!isSame(a,b)) return false;
        if(a == null) return true;
        if(!check(a.left, b.right)) return false;
        if(!check(a.right, b.left)) return false;
        return true;
    }
    
    public boolean isSymmetric(TreeNode root) {
        if(root == null) return true;
        return check(root.left, root.right);
    }
}
