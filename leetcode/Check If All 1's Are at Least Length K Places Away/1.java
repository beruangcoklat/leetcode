class Solution {
    public boolean kLengthApart(int[] nums, int k) {
        for(int i=0 ; i<nums.length ; i++){
            if(nums[i] == 0) continue;
            
            for(int j=0 ; j<k ; j++){
                if(i + 1 + j >= nums.length){
                    return true;
                }
                if(nums[i+ 1 + j] == 1){
                    return false;
                }
            }
        }
        return true;
    }
}
