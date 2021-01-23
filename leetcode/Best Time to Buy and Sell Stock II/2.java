class Solution {
    public int maxProfit(int[] prices) {
        int i = 0;
        int localMax = prices[0];
        int localMin = prices[0];
        int ret = 0;
        
        while(i + 1 < prices.length){
            while(i + 1 < prices.length && prices[i+1] < prices[i]){
                i++;
            }
            localMin = prices[i];
            while(i + 1 < prices.length && prices[i+1] > prices[i]){
                i++;
            }
            localMax = prices[i];
            ret += localMax - localMin;
            i++;
        }
        return ret;
    }
}
