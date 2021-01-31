class Solution {
    int[] arr;
    int[][] memo;
    boolean[][] visited;

    int solve(int idx, int n) {
        if (n == 0) {
            return 0;
        }

        if (visited[idx][n]) {
            return memo[idx][n];
        }

        int min = Integer.MAX_VALUE;
        for (int i = idx; i >= 0; i--) {
            int c = arr[i];
            if (c > n) continue;
            int curr = solve(i, n - arr[i]);
            if (curr != Integer.MAX_VALUE) {
                curr++;
            }
            if (curr < min) min = curr;
        }

        visited[idx][n] = true;
        memo[idx][n] = min;
        return min;
    }

    public int numSquares(int n) {
        memo = new int[101][10001];
        visited = new boolean[101][10001];
        arr = new int[101];
        int arrSize = 0;
        int i = 0;
        while (true) {
            i++;
            int x = i * i;
            if (x > n) {
                break;
            }
            arr[arrSize] = x;
            arrSize++;
        }

        return solve(arrSize - 1, n);
    }
}