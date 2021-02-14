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
    public List<List<Integer>> levelOrderBottom(TreeNode root) {
        List<List<Integer>> ret = new ArrayList<>();
        if (root == null) return ret;
        List<TreeNode> list = new ArrayList<>();
        list.add(root);
        while (!list.isEmpty()) {
            int size = list.size();
            List<Integer> row = new ArrayList<>();
            for (int aaa = 0; aaa < size; aaa++) {
                TreeNode curr = list.get(0);
                row.add(curr.val);
                list.remove(0);

                if (curr.left != null)
                    list.add(curr.left);
                if (curr.right != null)
                    list.add(curr.right);
            }
            ret.add(0, row);
        }

        return ret;
    }
}
