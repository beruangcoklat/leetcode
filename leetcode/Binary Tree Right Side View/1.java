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
    public List<Integer> rightSideView(TreeNode root) {
        List<Integer> ret = new ArrayList<>();

        if (root == null) return ret;

        List<TreeNode> list = new ArrayList<>();
        list.add(root);

        while (!list.isEmpty()) {
            int currSize = list.size();
            for (int i = 0; i < currSize; i++) {
                TreeNode curr = list.get(0);
                list.remove(0);
                if (i == 0) ret.add(curr.val);

                if (curr.right != null) list.add(curr.right);
                if (curr.left != null) list.add(curr.left);
            }
        }

        return ret;
    }
}