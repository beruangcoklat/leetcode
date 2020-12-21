class Solution {
    
    int getMax(int[] nums){
        int max = 0;
        for(int i=0 ; i<nums.length ; i++){
            if (i == 0 || nums[i] > max){
                max = nums[i];
            }
        }
        return max;
    }
    
    public List<Boolean> kidsWithCandies(int[] candies, int extraCandies) {
        int max = getMax(candies);
        List<Boolean> ret = new ArrayList();
        for(int i=0 ; i<candies.length ; i++){
            ret.add(candies[i] + extraCandies >= max);
        }
        
        return ret;
    }
}
