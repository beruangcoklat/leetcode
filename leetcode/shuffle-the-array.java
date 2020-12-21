class Solution {
    public int[] shuffle(int[] nums, int n) {
        int len = nums.length;
        int []ret = new int[nums.length];
        
        for(int i=0 ; i<nums.length / 2 ; i++){
            int idx = i * 2;
            ret[idx] = nums[i];
            ret[idx+1] = nums[i+n];
        }
        
        return ret;
    }
}
