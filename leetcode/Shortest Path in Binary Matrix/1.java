class Solution {
    class Point{
        int x,y;
        Point(int x, int y){
            this.x=x;
            this.y=y;
        }
    }

    public int shortestPathBinaryMatrix(int[][] grid) {
        if (grid[0][0] == 1) return -1;
        if (grid[grid.length - 1][grid[0].length - 1] == 1) return -1;

        int dirX[] = {0, 0, 1, -1, 1, 1, -1, -1};
        int dirY[] = {-1, 1, 0, 0, 1, -1, 1, -1};

        boolean[][] visited = new boolean[grid.length][grid[0].length];

        List<Point> list = new ArrayList<>();
        list.add(new Point(0, 0));

        for (int moves = 1; !list.isEmpty(); moves++) {
            int size = list.size();
            for (int aaa = 0; aaa < size; aaa++) {
                Point curr = list.get(0);
                list.remove(0);

                if (visited[curr.y][curr.x]) continue;
                visited[curr.y][curr.x] = true;

                if (curr.x == grid[0].length - 1 && curr.y == grid.length - 1) {
                    return moves;
                }

                for (int i = 0; i < dirX.length; i++) {
                    int nx = curr.x + dirX[i];
                    int ny = curr.y + dirY[i];
                    if (nx < 0 || nx >= grid[0].length || ny < 0 || ny >= grid.length) continue;
                    if (grid[ny][nx] == 1) continue;
                    list.add(new Point(nx, ny));
                }
            }
        }

        return -1;
    }
}
