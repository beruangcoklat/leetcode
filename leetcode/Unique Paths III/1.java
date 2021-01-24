class Solution {
    int nextX[] = {0, 0, 1, -1};
    int nextY[] = {1, -1, 0, 0};

    boolean[][] cloneVisited(boolean[][] visited) {
        boolean[][] ret = new boolean[visited.length][visited[0].length];
        for (int i = 0; i < visited.length; i++) {
            for (int j = 0; j < visited[0].length; j++) {
                ret[i][j] = visited[i][j];
            }
        }
        return ret;
    }

    int dfs(int x, int y, int m, int n, int[][] grid, boolean[][] visited, int path, int minimumPath) {
        if (grid[y][x] == 2 && path == minimumPath) {
            return 1;
        }

        int ret = 0;
        visited[y][x] = true;
        for (int i = 0; i < nextX.length; i++) {
            int nx = x + nextX[i];
            int ny = y + nextY[i];
            if (nx < 0 || nx >= n) continue;
            if (ny < 0 || ny >= m) continue;
            if (grid[ny][nx] == -1) continue;
            if (visited[ny][nx]) continue;
            ret += dfs(nx, ny, m, n, grid, cloneVisited(visited), path + 1, minimumPath);
        }
        
        return ret;
    }

    public int uniquePathsIII(int[][] grid) {
        int m = grid.length;
        int n = grid[0].length;

        int startX = 0, startY = 0;
        int totalObs = 0;

        for (int i = 0; i < m; i++) {
            for (int j = 0; j < n; j++) {
                int tile = grid[i][j];
                if (tile == 1) {
                    startX = j;
                    startY = i;
                } else if (tile == -1) {
                    totalObs++;
                }
            }
        }

        int minimumPath = (m * n) - 1 - totalObs;
        return dfs(startX, startY, m, n, grid, new boolean[m][n], 0, minimumPath);
    }
}