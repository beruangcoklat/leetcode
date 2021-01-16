import java.io.*;
import java.math.*;
import java.security.*;
import java.text.*;
import java.util.*;
import java.util.concurrent.*;
import java.util.regex.*;

public class Solution {

    // Complete the luckBalance function below.
    static int solve(int idx, int k, int[][] contests, HashMap<String, Integer> memo) {
        if (idx >= contests.length) {
            return 0;
        }

        int luck = contests[idx][0];
        int important = contests[idx][1];

        String key = idx + "-" + k;
        Integer mem = memo.get(key);
        if (mem != null) return mem;

        int notUse = solve(idx + 1, k, contests, memo) - luck;
        int max = notUse;
        if (k >= important) {
            int use = solve(idx + 1, k - important, contests, memo) + luck;
            if (use > notUse) {
                max = use;
            }
        }

        memo.put(key, max);
        return max;
    }

    static int luckBalance(int k, int[][] contests) {
        return solve(0, k, contests, new HashMap<String, Integer>());
    }

    private static final Scanner scanner = new Scanner(System.in);

    public static void main(String[] args) throws IOException {
        BufferedWriter bufferedWriter = new BufferedWriter(new FileWriter(System.getenv("OUTPUT_PATH")));

        String[] nk = scanner.nextLine().split(" ");

        int n = Integer.parseInt(nk[0]);

        int k = Integer.parseInt(nk[1]);

        int[][] contests = new int[n][2];

        for (int i = 0; i < n; i++) {
            String[] contestsRowItems = scanner.nextLine().split(" ");
            scanner.skip("(\r\n|[\n\r\u2028\u2029\u0085])?");

            for (int j = 0; j < 2; j++) {
                int contestsItem = Integer.parseInt(contestsRowItems[j]);
                contests[i][j] = contestsItem;
            }
        }

        int result = luckBalance(k, contests);

        bufferedWriter.write(String.valueOf(result));
        bufferedWriter.newLine();

        bufferedWriter.close();

        scanner.close();
    }
}
