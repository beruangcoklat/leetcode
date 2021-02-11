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
    public boolean isCousins(TreeNode root, int x, int y) {
        if (root == null) return false;

        List<TreeNode> list = new ArrayList<>();
        list.add(root);
        HashMap<Integer, Integer> parentMap = new HashMap<>();
        while (list.size() > 0) {
            int size = list.size();
            int counter = 0;

            for (int aaa = 0; aaa < size; aaa++) {
                TreeNode curr = list.get(0);
                list.remove(0);

                if (curr.val == x || curr.val == y) {
                    counter++;
                }

                if (counter == 2) {
                    return parentMap.get(x) != parentMap.get(y);
                }

                if (curr.right != null) {
                    list.add(curr.right);
                    parentMap.put(curr.right.val, curr.val);
                }
                if (curr.left != null) {
                    list.add(curr.left);
                    parentMap.put(curr.left.val, curr.val);
                }
            }

            if (counter > 0) {
                return false;
            }
        }
        return false;
    }
}
