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
    
    TreeNode root;
    
    TreeNode insert(TreeNode root, int val){
        if(root == null) return new TreeNode(val);
        if(val > root.val) root.right = insert(root.right, val);
        else root.left = insert(root.left, val);
        return root;
    }
    
    void preorder(TreeNode node, int low, int high){
        if(node == null) return;
        if(node.val >= low && node.val <= high)
            root = insert(root, node.val);
        preorder(node.left, low, high);
        preorder(node.right, low, high);
    }
    
    public TreeNode trimBST(TreeNode node, int low, int high) {
        root = null;
        preorder(node, low, high);
        return root;
    }
}