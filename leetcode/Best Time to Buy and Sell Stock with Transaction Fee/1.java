class Solution {    
    HashMap<String, Integer> memo;
    
    int solve(int idx, int[] prices, int state, int fee) {
        if (idx >= prices.length) {
            return 0;
        }

        String key = state + "-" + idx;
        Integer mem = memo.get(key);
        if (mem != null) {
            return mem;
        }

        int wait = solve(idx + 1, prices, state, fee);
        int max = wait;
        
        if (state == 0) {
            int buy = solve(idx + 1, prices, state + 1, fee) - prices[idx];
            if(buy > max) max = buy;
        } else if (state == 1) {
            int sell = solve(idx + 1, prices, state - 1, fee) + prices[idx] - fee;
            if(sell > max) max = sell;
        }

        memo.put(key, max);
        return max;
    }

    public int maxProfit(int[] prices, int fee) {
        memo = new HashMap();
        return solve(0, prices, 0, fee);
    }
}
