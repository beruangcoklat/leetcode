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
    boolean first;
    int prev, minDiff;

    void inorder(TreeNode root){
        if(root == null) return;
        inorder(root.left);

        if(first){
            first = false;
            prev = root.val;
        }else{
            int diff = root.val - prev;
            if(diff < minDiff) minDiff = diff;
            prev = root.val;
        }

        inorder(root.right);
    }

    public int minDiffInBST(TreeNode root) {
        minDiff = Integer.MAX_VALUE;
        first = true;
        inorder(root);
        return minDiff;
    }
}