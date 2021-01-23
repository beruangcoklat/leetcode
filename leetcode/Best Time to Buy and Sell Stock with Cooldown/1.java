class Solution {
    
    HashMap<String, Integer> memo;
    
    int solve(int idx, int[] prices, int state) {
        if (idx >= prices.length) {
            return 0;
        }

        String key = state + "-" + idx;
        Integer mem = memo.get(key);
        if (mem != null) {
            return mem;
        }

        int wait = solve(idx + 1, prices, state);
        int max = wait;
        
        if (state == 0) {
            int buy = solve(idx + 1, prices, state + 1) - prices[idx];
            if(buy > max) max = buy;
        } else if (state == 1) {
            int sell = solve(idx + 2, prices, state - 1) + prices[idx];
            if(sell > max) max = sell;
        }

        memo.put(key, max);
        return max;
    }

    public int maxProfit(int[] prices) {
        memo = new HashMap();
        return solve(0, prices, 0);
    }
}
