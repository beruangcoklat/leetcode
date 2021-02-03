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
    public int rangeSumBST(TreeNode root, int low, int high) {
        if(root == null) return 0;
        
        int sum = 0;
        List<TreeNode> list = new ArrayList();
        list.add(root);
        while(list.size() > 0){
            TreeNode curr = list.get(0);
            list.remove(0);
            
            if(curr == null) continue;
            if(curr.val >= low && curr.val <= high)
                sum += curr.val;
            
            if(curr.val <= high) list.add(curr.right);
            if(curr.val >= low) list.add(curr.left);
        }
        
        
        return sum;
    }
}