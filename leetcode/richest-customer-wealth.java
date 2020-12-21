class Solution {
    
    int getTotal(int[] nums){
        int total = 0;
        for(int i=0 ; i<nums.length ; i++){
            total += nums[i];    
        }
        return total;
    }
    
    public int maximumWealth(int[][] accounts) {
        int max = 0;
        for(int i=0 ; i<accounts.length ; i++){
            int localMax = getTotal(accounts[i]);
            if(i == 0 || localMax > max){
                max = localMax;
            }
        }
        return max;
    }
}
