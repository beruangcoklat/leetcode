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
    public List<List<Integer>> levelOrder(TreeNode root) {
        List<List<Integer>> ret = new ArrayList<>();
        if(root == null) return ret;
        
        List<TreeNode> list = new ArrayList<>();
        list.add(root);
        while (list.size() > 0) {
            int listSize = list.size();
            List<Integer> currLevel = new ArrayList<>();
            for (int i = 0; i < listSize; i++) {
                TreeNode curr = list.get(0);
                list.remove(0);
                currLevel.add(curr.val);
                if (curr.left != null) {
                    list.add(curr.left);
                }
                if (curr.right != null) {
                    list.add(curr.right);
                }
            }
            ret.add(currLevel);
        }
        return ret;
    }
}