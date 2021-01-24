class Solution {
    HashMap<String, Boolean> memo;

    boolean solve(boolean isAlice, int n) {
        if (n == 2) {
            return isAlice;
        }

        String key = isAlice + "-" + n;
        Boolean mem = memo.get(key);
        if (mem != null) return mem;

        for (int i = 1; i < n; i++) {
            if (n % i == 0) {
                boolean tmp = solve(!isAlice, n - i);
                if ((tmp && isAlice) || (!tmp && !isAlice)) {
                    memo.put(key, tmp);
                    return tmp;
                }
            }
        }

        memo.put(key, !isAlice);
        return !isAlice;
    }

    public boolean divisorGame(int N) {
        memo = new HashMap<>();
        return solve(true, N);
    }
}
