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
import java.math.BigDecimal;
import java.math.MathContext;

class Solution {
    public List<Double> averageOfLevels(TreeNode root) {
        List<Double> ret = new ArrayList<>();
        if (root == null) return ret;
        List<TreeNode> list = new ArrayList<>();
        list.add(root);
        while (!list.isEmpty()) {
            int size = list.size();
            BigDecimal total = new BigDecimal(0);
            for (int aaa = 0; aaa < size; aaa++) {
                TreeNode curr = list.get(0);
                total = total.add(new BigDecimal(curr.val));
                list.remove(0);

                if (curr.left != null)
                    list.add(curr.left);
                if (curr.right != null)
                    list.add(curr.right);
            }
            ret.add(total.divide(new BigDecimal(size), MathContext.DECIMAL128).doubleValue());
        }

        return ret;
    }
}
