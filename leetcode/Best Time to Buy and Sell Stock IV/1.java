class Solution {
    HashMap<String, Integer> memo = new HashMap<>();

    int solve(int idx, int k, int[] prices, int state, int tx) {
        if (idx == prices.length || tx == k) {
            return 0;
        }

        String key = state + "-" + idx + "-" + tx;
        Integer mem = memo.get(key);
        if (mem != null) {
            return mem;
        }

        int wait = solve(idx + 1, k, prices, state, tx);
        int max = wait;

        if (state == 0) {
            int buy = solve(idx + 1, k, prices, state + 1, tx);
            buy -= prices[idx];
            max = Math.max(buy, wait);
        } else if (state == 1) {
            int sell = solve(idx + 1, k, prices, state - 1, tx + 1);
            sell += prices[idx];
            max = Math.max(sell, wait);
        }

        memo.put(key, max);
        return max;
    }

    public int maxProfit(int k, int[] prices) {
        memo = new HashMap<>();
        return solve(0, k, prices, 0, 0);
    }
}
