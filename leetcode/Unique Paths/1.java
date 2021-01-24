class Solution {
    
    HashMap<String, Integer> memo;

    int solve(int x, int y, int m, int n) {
        if (x == n - 1 && y == m - 1) {
            return 1;
        }

        String key = String.format("%d-%d", x, y);
        Integer mem = memo.get(key);
        if (mem != null) {
            return mem;
        }

        int ret = 0;
        if (x + 1 < n) {
            ret += solve(x + 1, y, m, n);
        }
        if (y + 1 < m) {
            ret += solve(x, y + 1, m, n);
        }

        memo.put(key, ret);
        return ret;
    }

    public int uniquePaths(int m, int n) {
        memo = new HashMap<String, Integer>();
        return solve(0, 0, m, n);
    }
}