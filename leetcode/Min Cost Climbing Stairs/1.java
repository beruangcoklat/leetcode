class Solution {
    HashMap<Integer, Integer> memo;

    int solve(int idx, int[] cost) {
        if (idx >= cost.length) return 0;

        Integer mem = memo.get(idx);
        if (mem != null) return mem;

        int step1 = solve(idx + 1, cost) + cost[idx];
        int step2 = solve(idx + 2, cost) + cost[idx];

        int min = Math.min(step1, step2);
        memo.put(idx, min);
        return min;
    }

    public int minCostClimbingStairs(int[] cost) {
        memo = new HashMap<>();
        int startFrom0 = solve(0, cost);
        int startFrom1 = solve(1, cost);
        return Math.min(startFrom0, startFrom1);
    }
}
