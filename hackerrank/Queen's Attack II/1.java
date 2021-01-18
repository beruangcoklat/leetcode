import java.io.*;
import java.math.*;
import java.security.*;
import java.text.*;
import java.util.*;
import java.util.concurrent.*;
import java.util.regex.*;

public class Solution {

    static class Grid {
        int row;
        int col;

        Grid(int row, int col) {
            this.row = row;
            this.col = col;
        }
    }

    static HashMap<String, Byte> getObstaclesMap(int[][] obstacles) {
        HashMap<String, Byte> map = new HashMap<>();
        for (int[] obs : obstacles) {
            String key = (obs[0] - 1) + "-" + (obs[1] - 1);
            map.put(key, null);
        }
        return map;
    }

    static List<Grid> getGrids() {
        int[] nextX = {1, 1, 1, -1, -1, -1, 0, 0};
        int[] nextY = {0, 1, -1, 0, 1, -1, -1, 1};
        List<Grid> grids = new ArrayList<>();

        for (int i = 0; i < nextX.length; i++) {
            grids.add(new Grid(nextY[i], nextX[i]));
        }

        return grids;
    }

    static int queensAttack(int n, int k, int r_q, int c_q, int[][] obstacles) {
        r_q--;
        c_q--;

        HashMap<String, Byte> obsMap = getObstaclesMap(obstacles);
        List<Grid> grids = getGrids();

        int ret = 0;
        int count = 0;
        int mult = 0;

        while (true) {
            count = 0;
            mult++;

            for (int i = 0; i < grids.size(); i++) {
                int ny = (grids.get(i).row * mult) + r_q;
                int nx = (grids.get(i).col * mult) + c_q;
                
                if (nx < 0 || nx >= n || ny < 0 || ny >= n || obsMap.containsKey(ny + "-" + nx)) {
                    grids.remove(i);
                    i--;
                    continue;
                }
                
                count++;
            }

            if (count == 0) break;
            ret += count;
        }
        return ret;
    }

    private static final Scanner scanner = new Scanner(System.in);

    public static void main(String[] args) throws IOException {
        BufferedWriter bufferedWriter = new BufferedWriter(new FileWriter(System.getenv("OUTPUT_PATH")));

        String[] nk = scanner.nextLine().split(" ");

        int n = Integer.parseInt(nk[0]);

        int k = Integer.parseInt(nk[1]);

        String[] r_qC_q = scanner.nextLine().split(" ");

        int r_q = Integer.parseInt(r_qC_q[0]);

        int c_q = Integer.parseInt(r_qC_q[1]);

        int[][] obstacles = new int[k][2];

        for (int i = 0; i < k; i++) {
            String[] obstaclesRowItems = scanner.nextLine().split(" ");
            scanner.skip("(\r\n|[\n\r\u2028\u2029\u0085])?");

            for (int j = 0; j < 2; j++) {
                int obstaclesItem = Integer.parseInt(obstaclesRowItems[j]);
                obstacles[i][j] = obstaclesItem;
            }
        }

        int result = queensAttack(n, k, r_q, c_q, obstacles);

        bufferedWriter.write(String.valueOf(result));
        bufferedWriter.newLine();

        bufferedWriter.close();

        scanner.close();
    }
}
