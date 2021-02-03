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

    int prev;
    boolean first;
    boolean res;
    
    void inorder(TreeNode root){
        if(root == null) return;
        inorder(root.left);
        if(first){
            first = false;
            prev = root.val;
        }else{
            if(root.val <= prev){
                res = false;
                return;
            }
            prev = root.val;
        }
        inorder(root.right);
    }
    
    public boolean isValidBST(TreeNode root) {
        first = true;
        res = true;
        inorder(root);
        return res;
    }
}