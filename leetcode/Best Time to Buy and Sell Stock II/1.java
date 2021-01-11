class Solution {
    HashMap<String, Integer> memo = new HashMap<>();

    int solve(int idx, int[] prices, int state) {
        if (idx == prices.length) {
            return 0;
        }

        String key = state + "-" + idx;
        if (memo.containsKey(key)) {
            return memo.get(key);
        }

        int max = 0;
        if (state == 0) {
            int buy = solve(idx + 1, prices, 1) - prices[idx];
            int wait = solve(idx + 1, prices, 0);
            max = Math.max(buy, wait);
        } else if (state == 1) {
            int sell = solve(idx + 1, prices, 0) + prices[idx];
            int wait = solve(idx + 1, prices, 1);
            max = Math.max(sell, wait);
        }

        memo.put(key, max);
        return max;
    }

    public int maxProfit(int[] prices) {
        memo.clear();
        return solve(0, prices, 0);
    }
}
