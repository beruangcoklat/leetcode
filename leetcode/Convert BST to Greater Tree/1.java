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
    
    List<Integer> arr;
    int a[];
    int idx;
    
    void inorder(TreeNode root, boolean incr){
        if(root == null) return;
        inorder(root.left, incr);
        if(!incr){
            arr.add(root.val);
        }else{
            root.val += a[idx];
            idx++;
        }
        inorder(root.right, incr);
    }
    
    public TreeNode convertBST(TreeNode root) {
        arr = new ArrayList();
        inorder(root,false);
        idx = 0;
        a = new int[arr.size()];
        for(int i=arr.size()-2 ; i>=0 ; i--){
            a[i] = a[i+1] + arr.get(i+1);
        }
        inorder(root, true);
        
        return root;
    }
}
