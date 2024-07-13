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
    public List<List<Integer>> zigzagLevelOrder(TreeNode root) {
        List<TreeNode> list = new ArrayList<>();
        list.add(root);

        List<List<Integer>> result = new ArrayList<>();

        for (int level = 0; !list.isEmpty(); level++) {
            List<Integer> nums = new ArrayList<>();
            int size = list.size();
            for (int aaa = 0; aaa < size; aaa++) {
                TreeNode curr = list.get(0);
                list.remove(0);
                if (curr == null) {
                    continue;
                }
                nums.add(curr.val);
                list.add(curr.left);
                list.add(curr.right);
            }

            if (nums.isEmpty()) {
                continue;
            }

            if (level % 2 == 1) {
                for (int i = 0; i < nums.size() / 2; i++) {
                    int a = nums.get(i);
                    int b = nums.get(nums.size() - 1 - i);
                    nums.set(i, b);
                    nums.set(nums.size() - 1 - i, a);
                }
            }
            result.add(nums);
        }
        return result;
    }
}