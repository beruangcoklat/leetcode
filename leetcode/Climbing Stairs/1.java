class Solution {
    
    int []memo = new int[46];   
    
    int solve(int n) {
        if (n == 0) return 1;
        if (n < 0) return 0;
        if (memo[n] != 0) return memo[n];
        memo[n] = solve(n - 1) + solve(n - 2);
        return memo[n];
    }
    
    public int climbStairs(int n) {
        for(int i=0 ; i<n+1 ; i++) {
            memo[i] = 0;
        }
        return solve(n);
    }
}
