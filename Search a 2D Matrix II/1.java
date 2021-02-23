class Solution {
    public boolean searchMatrix(int[][] matrix, int target) {
        int row = matrix.length;
        int col = matrix[0].length;

        int x = 0;
        int y = 0;

        while (x + 1 < col && matrix[0][x] <= target) x++;
        while (y + 1 < row && matrix[y][0] <= target) y++;

        for (int yy = 0; yy <= y; yy++) {
            for (int xx = 0; xx <= x; xx++) {
                if (matrix[yy][xx] == target) {
                    return true;
                }
            }
        }
        return false;
    }
}
