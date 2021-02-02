import java.io.*;
import java.math.*;
import java.security.*;
import java.text.*;
import java.util.*;
import java.util.concurrent.*;
import java.util.regex.*;

public class Solution {

    static int nextX[] = {0, 0, 1, -1, 1, 1, -1, -1};
    static int nextY[] = {1, -1, 0, 0, 1, -1, 1, -1};

    static int dfs(int[][] matrix, int x, int y, boolean[][] visited) {
        visited[y][x] = true;

        int ret = 1;
        for (int i = 0; i < nextX.length; i++) {
            int nx = x + nextX[i];
            int ny = y + nextY[i];
            if (nx < 0 || nx >= matrix[0].length || ny < 0 || ny >= matrix.length) continue;
            if (visited[ny][nx]) continue;
            if (matrix[ny][nx] == 0) continue;

            ret += dfs(matrix, nx, ny, visited);
        }

        return ret;
    }

    static int connectedCell(int[][] matrix) {
        boolean[][] visited = new boolean[matrix.length][matrix[0].length];
        int max = 0;
        for (int i = 0; i < matrix.length; i++) {
            for (int j = 0; j < matrix[i].length; j++) {
                if (matrix[i][j] == 0) continue;
                if (visited[i][j]) continue;
                int curr = dfs(matrix, j, i, visited);
                max = Math.max(max, curr);
            }
        }
        return max;
    }

    private static final Scanner scanner = new Scanner(System.in);

    public static void main(String[] args) throws IOException {
        BufferedWriter bufferedWriter = new BufferedWriter(new FileWriter(System.getenv("OUTPUT_PATH")));

        int n = scanner.nextInt();
        scanner.skip("(\r\n|[\n\r\u2028\u2029\u0085])?");

        int m = scanner.nextInt();
        scanner.skip("(\r\n|[\n\r\u2028\u2029\u0085])?");

        int[][] matrix = new int[n][m];

        for (int i = 0; i < n; i++) {
            String[] matrixRowItems = scanner.nextLine().split(" ");
            scanner.skip("(\r\n|[\n\r\u2028\u2029\u0085])?");

            for (int j = 0; j < m; j++) {
                int matrixItem = Integer.parseInt(matrixRowItems[j]);
                matrix[i][j] = matrixItem;
            }
        }

        int result = connectedCell(matrix);

        bufferedWriter.write(String.valueOf(result));
        bufferedWriter.newLine();

        bufferedWriter.close();

        scanner.close();
    }
}
