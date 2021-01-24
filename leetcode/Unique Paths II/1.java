class Solution {
    HashMap<String, Integer> memo;

    int solve(int x, int y, int m, int n, int[][] obstacleGrid) {
        if (x == n - 1 && y == m - 1) {
            return obstacleGrid[y][x] == 0 ? 1 : 0;
        }

        String key = String.format("%d-%d", x, y);
        Integer mem = memo.get(key);
        if (mem != null) {
            return mem;
        }

        int ret = 0;
        if (x + 1 < n && obstacleGrid[y][x + 1] == 0) {
            ret += solve(x + 1, y, m, n, obstacleGrid);
        }
        if (y + 1 < m && obstacleGrid[y + 1][x] == 0) {
            ret += solve(x, y + 1, m, n, obstacleGrid);
        }

        memo.put(key, ret);
        return ret;
    }

    public int uniquePathsWithObstacles(int[][] obstacleGrid) {
        if (obstacleGrid[0][0] == 1) {
            return 0;
        }
        memo = new HashMap<String, Integer>();
        int m = obstacleGrid.length;
        int n = obstacleGrid[0].length;
        return solve(0, 0, m, n, obstacleGrid);
    }
}