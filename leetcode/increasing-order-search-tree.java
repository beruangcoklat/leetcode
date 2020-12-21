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
    ArrayList<Integer> arr = new ArrayList();
    
    void inorder(TreeNode node){
        if(node == null) return;
        inorder(node.left);
        arr.add(node.val);
        inorder(node.right);
    }
    
    public TreeNode increasingBST(TreeNode root) {
        inorder(root);
        TreeNode newRoot = null;
        TreeNode curr = null;
        for(int i=0 ; i<arr.size() ; i++){
            if(newRoot == null){
                newRoot = new TreeNode(arr.get(i));
                curr = newRoot;
                continue;
            }
            
            curr.right = new TreeNode(arr.get(i));
            curr = curr.right;
        }
        return newRoot;
    }
}
