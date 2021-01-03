class Solution {
    
    int [][]memo;
    
    int solve(int x, int y, int m, int n) {
        if (x == n - 1 && y == m - 1) {
            return 1;
        }
        
        if(memo[x][y] != 0){
            return memo[x][y];
        }

        boolean right = x + 1 < n;
        boolean down = y + 1 < m;

        int ret = 0;
        if (right) {
            ret += solve(x + 1, y, m, n);
        }
        if (down) {
            ret += solve(x, y + 1, m, n);
        }

        memo[x][y] = ret;
        return ret;
    }
    
    public int uniquePaths(int m, int n) {
        memo = new int[100][100];
        return solve(0, 0, m, n);
    }
}
