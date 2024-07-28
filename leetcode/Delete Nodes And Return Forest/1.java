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
    Map<Integer, Void> deleteMap;
    List<TreeNode> result;

    public List<TreeNode> delNodes(TreeNode root, int[] to_delete) {
        result = new ArrayList<>();
        deleteMap = new HashMap<>();

        for (int a : to_delete) {
            deleteMap.put(a, null);
        }

        if (!deleteMap.containsKey(root.val)) {
            result.add(root);
        }

        dfs(root, null, true);
        return result;
    }

    void dfs(TreeNode root, TreeNode parent, boolean isLeftChild) {
        if (root == null) return;

        TreeNode left = root.left;
        TreeNode right = root.right;

        if (deleteMap.containsKey(root.val)) {
            if (root.left != null && !deleteMap.containsKey(root.left.val)) {
                result.add(root.left);
            }
            if (root.right != null && !deleteMap.containsKey(root.right.val)) {
                result.add(root.right);
            }
            root.left = null;
            root.right = null;

            if (parent != null) {
                if (isLeftChild) {
                    parent.left = null;
                } else {
                    parent.right = null;
                }
            }
        }
        
        dfs(left, root, true);
        dfs(right, root, false);
    }
}