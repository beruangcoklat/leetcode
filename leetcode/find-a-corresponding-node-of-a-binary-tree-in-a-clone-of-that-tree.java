/**
 * Definition for a binary tree node.
 * public class TreeNode {
 *     int val;
 *     TreeNode left;
 *     TreeNode right;
 *     TreeNode(int x) { val = x; }
 * }
 */

class Solution {
    
    class Result{
        TreeNode ori, kw;
    }
    
    public final Result solve(final TreeNode original, final TreeNode cloned, final TreeNode target){
        if(original == target){
            Result res = new Result();
            res.ori = original;
            res.kw = cloned;
            return res;
        }
        
        if(original.right != null){
            Result curr = solve(original.right, cloned.right, target);
            if(curr != null && curr.ori == target){
                return curr;
            }
        }
        
        if(original.left != null){
            Result curr = solve(original.left, cloned.left, target);
            if(curr != null && curr.ori == target){
                return curr;
            }
        }
        
        return null;
    }
    
    public final TreeNode getTargetCopy(final TreeNode original, final TreeNode cloned, final TreeNode target) {
        Result res = solve(original, cloned, target);
        return res.kw;
    }
}
